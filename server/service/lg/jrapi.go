package lg

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrclientrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrclientresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrresponse"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	lg2 "github.com/flipped-aurora/gin-vue-admin/server/utils/lg"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"net/http"
	"strconv"
	"time"
)

type JRAPIService struct {
}

func (jrAPIService *JRAPIService) ApplyOrder(reApply jrrequest.JRAPIApply) (resApply jrresponse.JRAPIApply, err error) {
	if reApply.OrderNo == nil ||
		reApply.ApplyNo == nil ||
		reApply.ProductNo == nil ||
		reApply.ProductType == nil ||
		reApply.ProductRate == nil ||
		reApply.ElogAmount == nil ||
		reApply.ProjectGuid == nil ||
		reApply.ProjectName == nil ||
		reApply.ProjectNo == nil ||
		reApply.TenderDeposit == nil ||
		reApply.DepositStartDate == nil ||
		reApply.DepositEndDate == nil ||
		reApply.OpenBeginDate == nil ||
		reApply.ElogTemplateNo == nil ||
		reApply.ElogTemplateName == nil ||
		reApply.InsuredName == nil ||
		reApply.InsuredCreditCode == nil ||
		reApply.InsureName == nil ||
		reApply.InsureCreditCode == nil ||
		reApply.InsureLegalName == nil ||
		reApply.InsureLegalIdCard == nil ||
		reApply.InsureAddress == nil ||
		reApply.ApplicantName == nil ||
		reApply.ApplicantIdCard == nil ||
		reApply.ApplicantTel == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if !errors.Is(tx.Where("order_no = ? AND apply_no = ?", reApply.OrderNo, reApply.ApplyNo).First(&lg.Apply{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("相同订单和开函申请已经存在")
		}

		isRepushed := false
		isOfflineRefund := false
		order := &lg.Order{
			OrderNo:         reApply.OrderNo,
			IsRepushed:      &isRepushed,
			IsOfflineRefund: &isOfflineRefund,
		}

		var employee sysModel.SysUser
		if reApply.ApplicantAuthCode != nil {
			if !errors.Is(tx.Where("employee_no = ?", reApply.ApplicantAuthCode).First(&employee).Error, gorm.ErrRecordNotFound) {
				order.EmployeeID = &employee.ID
			}
		}

		if err = tx.Create(&order).Error; err != nil {
			return errors.New("创建订单失败")
		}
		// 匹配项目，进入自动化流程
		auditStatus := int64(1)
		auditOpinion := "待受理"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		realElogAmount := math.Trunc(*reApply.TenderDeposit*global.GVA_CONFIG.Insurance.ElogRate*1e2+0.5) * 1e-2
		realElogRate := global.GVA_CONFIG.Insurance.ElogRate
		tenderDeposit := *reApply.TenderDeposit
		insuranceName := global.GVA_CONFIG.Insurance.Name
		insuranceCreditCode := global.GVA_CONFIG.Insurance.CreditCode
		resApply.OrderNo = reApply.OrderNo
		resApply.ApplyNo = reApply.ApplyNo
		var apply lg.Apply
		var project lg.Project
		err = tx.Model(&lg.Project{}).Where("project_no = ? AND is_enable = TRUE", reApply.ProjectNo).Order("created_at desc").First(&project).Error
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("连接项目数据库失败")
			} else {
				auditStatus = int64(3)
				auditOpinion = "未匹配该项目"
				resApply.AuditStatus = &auditStatus
				resApply.AuditOpinion = &auditOpinion
				resApply.AuditDate = &auditDate
			}
		} else {
			order.ProjectID = &project.ID
			if err = tx.Save(&order).Error; err != nil {
				return errors.New("更新订单项目失败")
			}
			if *project.IsAutoMatic {
				auditStatus = int64(2)
				auditOpinion = "受理成功"
				realElogAmount = math.Trunc(*project.TenderDeposit*global.GVA_CONFIG.Insurance.ElogRate*1e2+0.5) * 1e-2
				if realElogAmount < global.GVA_CONFIG.Insurance.ElogMinAmount {
					realElogAmount = global.GVA_CONFIG.Insurance.ElogMinAmount
				}
				tenderDeposit = *project.TenderDeposit
				resApply.AuditStatus = &auditStatus
				resApply.AuditOpinion = &auditOpinion
				resApply.AuditDate = &auditDate
				resApply.RealElogAmount = &realElogAmount
				resApply.RealElogRate = &realElogRate
				resApply.TenderDeposit = &tenderDeposit
				resApply.InsuranceName = &insuranceName
				resApply.InsuranceCreditCode = &insuranceCreditCode
				apply.RealElogAmount = &realElogAmount
				apply.RealElogRate = &realElogRate
				apply.InsuranceName = &insuranceName
				apply.InsuranceCreditCode = &insuranceCreditCode
			} else {
				resApply.AuditStatus = &auditStatus
				resApply.AuditOpinion = &auditOpinion
				resApply.AuditDate = &auditDate
			}
		}
		attachInfo, _ := json.Marshal(reApply.AttachInfo)
		attachInfoString := string(attachInfo)
		productType, _ := strconv.ParseInt(*reApply.ProductType, 10, 64)
		apply.OrderNo = reApply.OrderNo
		apply.ApplyNo = reApply.ApplyNo
		apply.OrderID = &order.ID
		apply.ProductNo = reApply.ProductNo
		apply.ProductType = &productType
		apply.ProductRate = reApply.ProductRate
		apply.ElogAmount = reApply.ElogAmount
		apply.ProjectGuid = reApply.ProjectGuid
		apply.ProjectName = reApply.ProjectName
		apply.ProjectNo = reApply.ProjectNo
		apply.TenderDeposit = reApply.TenderDeposit
		apply.DepositStartDate = reApply.DepositStartDate
		apply.DepositEndDate = reApply.DepositEndDate
		apply.OpenBeginDate = reApply.OpenBeginDate
		apply.ElogTemplateNo = reApply.ElogTemplateNo
		apply.ElogTemplateName = reApply.ElogTemplateName
		apply.InsuredName = reApply.InsuredName
		apply.InsuredCreditCode = reApply.InsuredCreditCode
		apply.InsuredAddress = reApply.InsuredAddress
		apply.InsureName = reApply.InsureName
		apply.InsureCreditCode = reApply.InsureCreditCode
		apply.InsureLegalName = reApply.InsureLegalName
		apply.InsureLegalIdCard = reApply.InsureLegalIdCard
		apply.InsureAddress = reApply.InsureAddress
		apply.ApplicantName = reApply.ApplicantName
		apply.ApplicantIdCard = reApply.ApplicantIdCard
		apply.ApplicantTel = reApply.ApplicantTel
		apply.ApplicantAuthCode = reApply.ApplicantAuthCode
		apply.AttachInfo = &attachInfoString
		apply.AuditStatus = &auditStatus
		apply.AuditOpinion = &auditOpinion
		apply.AuditDate = &auditDate
		if err = tx.Create(&apply).Error; err != nil {
			return errors.New("创建申请失败")
		}
		order.ApplyID = &apply.ID
		if err = tx.Save(&order).Error; err != nil {
			return errors.New("更新订单失败")
		}
		return nil
	})
	return
}

func (jrAPIService *JRAPIService) PayPush(rePayPush jrrequest.JRAPIPayPush) (resPayPush jrresponse.JRAPIPayPush, err error) {
	if rePayPush.OrderNo == nil ||
		rePayPush.PayNo == nil ||
		rePayPush.PayAmount == nil ||
		rePayPush.PayTime == nil ||
		rePayPush.PayTransNo == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	var order lg.Order
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if !errors.Is(tx.Where("order_no = ? AND pay_no = ?", rePayPush.OrderNo, rePayPush.PayNo).
			First(&lg.Pay{}).Error, gorm.ErrRecordNotFound) {
			receiveResult := "success"
			resPayPush = jrresponse.JRAPIPayPush{
				ReceiveResult: &receiveResult,
			}
			return nil
		}
		if err = tx.Where("order_no = ?", rePayPush.OrderNo).Preload(clause.Associations).Preload("Project.Template").First(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("该订单" + *order.OrderNo + "不存在")
			} else {
				return errors.New("查询相应订单" + *order.OrderNo + "失败")
			}
		}
		pay := &lg.Pay{
			OrderID:    &order.ID,
			OrderNo:    rePayPush.OrderNo,
			PayNo:      rePayPush.PayNo,
			PayAmount:  rePayPush.PayAmount,
			PayTime:    rePayPush.PayTime,
			PayTransNo: rePayPush.PayTransNo,
		}
		if err = tx.Create(&pay).Error; err != nil {
			return errors.New(*order.OrderNo + "创建支付结果失败")
		}
		order.PayID = &pay.ID
		if err = tx.Save(&order).Error; err != nil {
			return errors.New("更新" + *order.OrderNo + "订单失败")
		}
		receiveResult := "success"
		resPayPush = jrresponse.JRAPIPayPush{
			ReceiveResult: &receiveResult,
		}

		return nil
	})
	if err == nil {
		go openLetterSync(order)
	}
	return
}

func (jrAPIService *JRAPIService) QueryInfo(reQueryInfo jrrequest.JRAPIQueryInfo) (resQueryInfo jrresponse.JRAPIQueryInfo, err error) {
	if reQueryInfo.ElogNo == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var order lg.Order
		if err = tx.Model(&lg.Order{}).Joins("Letter").Where("Letter.elog_no = ?", reQueryInfo.ElogNo).Preload(clause.Associations).First(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("该电子保函" + *reQueryInfo.ElogNo + "不存在")
			} else {
				return errors.New("查询相应保函" + *reQueryInfo.ElogNo + "失败")
			}
		}
		elogAmount := *order.Pay.PayAmount
		insuranceName := *order.Letter.InsuranceName
		insuranceCreditCode := *order.Letter.InsuranceCreditCode
		elogOutDate := *order.Letter.ElogOutDate
		var elogUrl string
		var elogEncryptUrl string
		var tenderDeposit float64
		var insureStartDate string
		var insureEndDate string
		var insureDay int64
		var validateCode string
		if order.Delay != nil {
			elogUrl = global.GVA_CONFIG.Insurance.APIDomain + "/delayFileDownload?elog=" + *order.Delay.ElogUrl
			elogEncryptUrl = global.GVA_CONFIG.Insurance.APIDomain + "/delayFileDownload?elog=" + *order.Delay.ElogEncryptUrl + "&type=encrypt"
			tenderDeposit = *order.Delay.TenderDeposit
			insureStartDate = *order.Delay.InsureStartDate
			insureEndDate = *order.Delay.InsureEndDate
			insureDay = *order.Delay.InsureDay
			validateCode = *order.Delay.ValidateCode
		} else {
			elogUrl = global.GVA_CONFIG.Insurance.APIDomain + "/letterFileDownload?elog=" + *order.Letter.ElogUrl
			elogEncryptUrl = global.GVA_CONFIG.Insurance.APIDomain + "/letterFileDownload?elog=" + *order.Letter.ElogEncryptUrl + "&type=encrypt"
			tenderDeposit = *order.Letter.TenderDeposit
			insureStartDate = *order.Letter.InsureStartDate
			insureEndDate = *order.Letter.InsureEndDate
			insureDay = *order.Letter.InsureDay
			validateCode = *order.Delay.ValidateCode
		}
		resQueryInfo = jrresponse.JRAPIQueryInfo{
			OrderNo:             order.OrderNo,
			ElogNo:              order.Letter.ElogNo,
			ProductNo:           order.Apply.ProductNo,
			ProductType:         order.Apply.ProductType,
			ProductRate:         order.Apply.ProductRate,
			ElogAmount:          &elogAmount,
			InsuranceName:       &insuranceName,
			InsuranceCreditCode: &insuranceCreditCode,
			ElogOutDate:         &elogOutDate,
			ElogUrl:             &elogUrl,
			ElogEncryptUrl:      &elogEncryptUrl,
			TenderDeposit:       &tenderDeposit,
			InsureStartDate:     &insureStartDate,
			InsureEndDate:       &insureEndDate,
			InsureDay:           &insureDay,
			ValidateCode:        &validateCode,
		}
		return nil
	})
	return
}

func (jrAPIService *JRAPIService) RevokePush(reRevokePush jrrequest.JRAPIRevokePush) (resRevokePush jrresponse.JRAPIRevokePush, err error) {
	if reRevokePush.OrderNo == nil ||
		reRevokePush.RevokeOrigin == nil ||
		reRevokePush.RevokeReason == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if !errors.Is(tx.Where("order_no = ?", reRevokePush.OrderNo).
			First(&lg.Revoke{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("该订单" + *reRevokePush.OrderNo + "已经撤单")
		}
		var order lg.Order
		if err = tx.Where("order_no = ?", reRevokePush.OrderNo).First(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("该订单" + *order.OrderNo + "不存在")
			} else {
				return errors.New("查询相应订单" + *order.OrderNo + "失败")
			}
		}
		revoke := &lg.Revoke{
			OrderID:      &order.ID,
			OrderNo:      reRevokePush.OrderNo,
			RevokeOrigin: reRevokePush.RevokeOrigin,
			RevokeReason: reRevokePush.RevokeReason,
		}
		if err = tx.Create(&revoke).Error; err != nil {
			return errors.New("创建" + *order.OrderNo + "撤单失败")
		}
		order.RevokeID = &revoke.ID
		if err = tx.Save(&order).Error; err != nil {
			return errors.New("更新" + *order.OrderNo + "订单失败")
		}
		receiveResult := "success"
		resRevokePush = jrresponse.JRAPIRevokePush{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (jrAPIService *JRAPIService) ApplyDelay(reApplyDelay jrrequest.JRAPIApplyDelay) (resApplyDelay jrresponse.JRAPIApplyDelay, err error) {
	if reApplyDelay.OrderNo == nil ||
		reApplyDelay.ApplyNo == nil ||
		reApplyDelay.ElogNo == nil ||
		reApplyDelay.ProjectGuid == nil ||
		reApplyDelay.ProjectName == nil ||
		reApplyDelay.ProjectNo == nil ||
		reApplyDelay.TenderDeposit == nil ||
		reApplyDelay.DepositStartDate == nil ||
		reApplyDelay.DepositEndDate == nil ||
		reApplyDelay.OpenBeginDate == nil ||
		reApplyDelay.InsureName == nil ||
		reApplyDelay.InsureCreditCode == nil ||
		reApplyDelay.ApplicantName == nil ||
		reApplyDelay.ApplicantIdCard == nil ||
		reApplyDelay.ApplicantTel == nil ||
		reApplyDelay.Reason == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var order lg.Order
		if err = tx.Where("order_no = ?", reApplyDelay.OrderNo).First(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("该订单" + *order.OrderNo + "不存在")
			} else {
				return errors.New("查询相应订单" + *order.OrderNo + "失败")
			}
		}
		if err = tx.Where("order_no = ?", reApplyDelay.OrderNo).Delete(&lg.Delay{}).Error; err != nil {
			return errors.New("删除历史延期数据失败")
		}
		attachInfo, _ := json.Marshal(reApplyDelay.AttachInfo)
		attachInfoString := string(attachInfo)
		delay := &lg.Delay{
			OrderID:          &order.ID,
			OrderNo:          reApplyDelay.OrderNo,
			ApplyNo:          reApplyDelay.ApplyNo,
			ElogNo:           reApplyDelay.ElogNo,
			ProjectGuid:      reApplyDelay.ProjectGuid,
			ProjectName:      reApplyDelay.ProjectName,
			ProjectNo:        reApplyDelay.ProjectNo,
			TenderDeposit:    reApplyDelay.TenderDeposit,
			DepositStartDate: reApplyDelay.DepositStartDate,
			DepositEndDate:   reApplyDelay.DepositEndDate,
			OpenBeginDate:    reApplyDelay.OpenBeginDate,
			InsureName:       reApplyDelay.InsureName,
			InsureCreditCode: reApplyDelay.InsureCreditCode,
			ApplicantName:    reApplyDelay.ApplicantName,
			ApplicantIdCard:  reApplyDelay.ApplicantIdCard,
			ApplicantTel:     reApplyDelay.ApplicantTel,
			Reason:           reApplyDelay.Reason,
			AttachInfo:       &attachInfoString,
		}
		if err = tx.Create(&delay).Error; err != nil {
			return errors.New("创建" + *order.OrderNo + "延期申请失败")
		}
		order.DelayID = &delay.ID
		if err = tx.Save(&order).Error; err != nil {
			return errors.New("更新" + *order.OrderNo + "订单失败")
		}
		receiveResult := "success"
		resApplyDelay = jrresponse.JRAPIApplyDelay{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (jrAPIService *JRAPIService) ApplyRefund(reApplyRefund jrrequest.JRAPIApplyRefund) (resApplyRefund jrresponse.JRAPIApplyRefund, err error) {
	if reApplyRefund.OrderNo == nil ||
		reApplyRefund.ApplyNo == nil ||
		reApplyRefund.ElogNo == nil ||
		reApplyRefund.InsureName == nil ||
		reApplyRefund.InsureCreditCode == nil ||
		reApplyRefund.ApplicantName == nil ||
		reApplyRefund.ApplicantIdCard == nil ||
		reApplyRefund.ApplicantTel == nil ||
		reApplyRefund.Reason == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var order lg.Order
		if err = tx.Where("order_no = ?", reApplyRefund.OrderNo).First(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("该订单" + *order.OrderNo + "不存在")
			} else {
				return errors.New("查询相应订单" + *order.OrderNo + "失败")
			}
		}
		if err = tx.Where("order_no = ?", reApplyRefund.OrderNo).Delete(&lg.Refund{}).Error; err != nil {
			return errors.New("删除历史退函数据失败")
		}
		attachInfo, _ := json.Marshal(reApplyRefund.AttachInfo)
		attachInfoString := string(attachInfo)
		auditStatus := int64(1)
		auditOpinion := "待受理"
		refund := &lg.Refund{
			OrderID:          &order.ID,
			OrderNo:          reApplyRefund.OrderNo,
			ApplyNo:          reApplyRefund.ApplyNo,
			ElogNo:           reApplyRefund.ElogNo,
			InsureName:       reApplyRefund.InsureName,
			InsureCreditCode: reApplyRefund.InsureCreditCode,
			ApplicantName:    reApplyRefund.ApplicantName,
			ApplicantIdCard:  reApplyRefund.ApplicantIdCard,
			ApplicantTel:     reApplyRefund.ApplicantTel,
			Reason:           reApplyRefund.Reason,
			AttachInfo:       &attachInfoString,
			AuditStatus:      &auditStatus,
			AuditOpinion:     &auditOpinion,
		}
		if err = tx.Create(&refund).Error; err != nil {
			return errors.New("创建" + *order.OrderNo + "退函申请失败")
		}
		order.RefundID = &refund.ID
		if err = tx.Save(&order).Error; err != nil {
			return errors.New("更新" + *order.OrderNo + "订单失败")
		}
		receiveResult := "success"
		resApplyRefund = jrresponse.JRAPIApplyRefund{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (jrAPIService *JRAPIService) ApplyClaim(reApplyClaim jrrequest.JRAPIApplyClaim) (resApplyClaim jrresponse.JRAPIApplyClaim, err error) {
	if reApplyClaim.OrderNo == nil ||
		reApplyClaim.ApplyNo == nil ||
		reApplyClaim.ElogNo == nil ||
		reApplyClaim.InsuredName == nil ||
		reApplyClaim.InsuredCreditCode == nil ||
		reApplyClaim.InsuredBankNo == nil ||
		reApplyClaim.InsuredBankName == nil ||
		reApplyClaim.ApplicantName == nil ||
		reApplyClaim.ApplicantIdCard == nil ||
		reApplyClaim.ApplicantTel == nil ||
		reApplyClaim.ClaimAmount == nil ||
		reApplyClaim.Reason == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if !errors.Is(tx.Where("order_no = ?", reApplyClaim.OrderNo).
			First(&lg.Claim{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("该订单" + *reApplyClaim.OrderNo + "已有理赔申请")
		}
		var order lg.Order
		if err = tx.Where("order_no = ?", reApplyClaim.OrderNo).First(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("该订单" + *order.OrderNo + "不存在")
			} else {
				return errors.New("查询相应订单" + *order.OrderNo + "失败")
			}
		}
		if err = tx.Where("order_no = ?", reApplyClaim.OrderNo).Delete(&lg.Claim{}).Error; err != nil {
			return errors.New("删除历史理赔数据失败")
		}
		attachInfo, _ := json.Marshal(reApplyClaim.AttachInfo)
		attachInfoString := string(attachInfo)
		claim := &lg.Claim{
			OrderID:           &order.ID,
			OrderNo:           reApplyClaim.OrderNo,
			ApplyNo:           reApplyClaim.ApplyNo,
			ElogNo:            reApplyClaim.ElogNo,
			InsuredName:       reApplyClaim.InsuredName,
			InsuredCreditCode: reApplyClaim.InsuredCreditCode,
			InsuredBankNo:     reApplyClaim.InsuredBankNo,
			InsuredBankName:   reApplyClaim.InsuredBankName,
			ApplicantName:     reApplyClaim.ApplicantName,
			ApplicantIdCard:   reApplyClaim.ApplicantIdCard,
			ApplicantTel:      reApplyClaim.ApplicantTel,
			ClaimAmount:       reApplyClaim.ClaimAmount,
			Reason:            reApplyClaim.Reason,
			AttachInfo:        &attachInfoString,
		}
		if err = tx.Create(&claim).Error; err != nil {
			return errors.New("创建" + *order.OrderNo + "理赔申请失败")
		}
		order.ClaimID = &claim.ID
		if err = tx.Save(&order).Error; err != nil {
			return errors.New("更新" + *order.OrderNo + "订单失败")
		}
		receiveResult := "success"
		resApplyClaim = jrresponse.JRAPIApplyClaim{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (jrAPIService *JRAPIService) LogoutPush(reLogoutPush jrrequest.JRAPILogoutPush) (resLogoutPush jrresponse.JRAPILogoutPush, err error) {
	if reLogoutPush.ProjectGuid == nil ||
		reLogoutPush.ProjectName == nil ||
		reLogoutPush.ProjectNo == nil ||
		reLogoutPush.Reason == nil ||
		reLogoutPush.LogoutType == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if !errors.Is(tx.Where("project_no = ?", reLogoutPush.ProjectNo).
			First(&lg.Logout{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("该项目" + *reLogoutPush.ProjectNo + "已有销函通知")
		}
		logout := &lg.Logout{
			ProjectGuid:         reLogoutPush.ProjectGuid,
			ProjectName:         reLogoutPush.ProjectName,
			ProjectNo:           reLogoutPush.ProjectNo,
			Reason:              reLogoutPush.Reason,
			LogoutType:          reLogoutPush.LogoutType,
			WinBidderName:       reLogoutPush.WinBidderName,
			WinBidderCreditCode: reLogoutPush.WinBidderCreditCode,
		}
		if err = tx.Create(&logout).Error; err != nil {
			return errors.New("创建" + *logout.ProjectNo + "销函通知失败")
		}
		var orders []lg.Order
		if err = tx.Model(&lg.Order{}).Joins("Apply").Where("Apply.project_guid = ?", reLogoutPush.ProjectGuid).Find(&orders).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("查询对应项目" + *reLogoutPush.ProjectNo + "失败")
			}
		}
		if len(orders) > 0 {
			for i := range orders {
				orders[i].LogoutID = &logout.ID
			}
			err = tx.Save(&orders).Error
			if err != nil {
				return errors.New("更新" + *reLogoutPush.ProjectNo + "订单失败")
			}
		}
		receiveResult := "success"
		resLogoutPush = jrresponse.JRAPILogoutPush{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (jrAPIService *JRAPIService) ApplyInvoice(reApplyInvoice jrrequest.JRAPIApplyInvoice) (resApplyInvoice jrresponse.JRAPIApplyInvoice, err error) {
	if reApplyInvoice.ApplyNo == nil ||
		reApplyInvoice.InvoiceTotalAmount == nil ||
		reApplyInvoice.InvoiceType == nil ||
		reApplyInvoice.InvoiceTileType == nil ||
		reApplyInvoice.InvoiceTile == nil ||
		reApplyInvoice.OrderList == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if !errors.Is(tx.Where("apply_no = ?", reApplyInvoice.ApplyNo).
			First(&lg.InvoiceApply{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("该发票申请" + *reApplyInvoice.ApplyNo + "已经存在")
		}
		orderList, _ := json.Marshal(reApplyInvoice.OrderList)
		orderListString := string(orderList)
		invoiceApply := &lg.InvoiceApply{
			ApplyNo:            reApplyInvoice.ApplyNo,
			InvoiceTotalAmount: reApplyInvoice.InvoiceTotalAmount,
			InvoiceType:        reApplyInvoice.InvoiceType,
			InvoiceTileType:    reApplyInvoice.InvoiceTileType,
			InvoiceTile:        reApplyInvoice.InvoiceTile,
			TaxNo:              reApplyInvoice.TaxNo,
			BankName:           reApplyInvoice.BankName,
			BankNo:             reApplyInvoice.BankNo,
			CompanyAddress:     reApplyInvoice.CompanyAddress,
			CompanyTel:         reApplyInvoice.CompanyTel,
			Remarks:            reApplyInvoice.Remarks,
			OrderList:          &orderListString,
		}
		if err = tx.Create(&invoiceApply).Error; err != nil {
			return errors.New("创建" + *reApplyInvoice.ApplyNo + "发票申请失败")
		}
		receiveResult := "success"
		resApplyInvoice = jrresponse.JRAPIApplyInvoice{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (jrAPIService *JRAPIService) LetterFileDownload(elog string, encrypt bool) (file lg.File, err error) {
	var letter lg.Letter
	db := global.GVA_DB.Model(&lg.Letter{})
	if encrypt {
		db = db.Where("elog_encrypt_url = ?", elog)
	} else {
		db = db.Where("elog_url = ?", elog)
	}
	err = db.Preload("Order").Preload("Order.Delay").Preload("ElogFile").Preload("ElogEncryptFile").Order("created_at desc").First(&letter).Error
	if err != nil {
		return lg.File{}, err
	}
	if letter.Order.DelayID != nil && *letter.Order.Delay.AuditStatus == 2 {
		return lg.File{}, err
	}
	if encrypt {
		file = *letter.ElogEncryptFile
	} else {
		file = *letter.ElogFile
	}
	return
}

func (jrAPIService *JRAPIService) DelayFileDownload(elog string, encrypt bool) (file lg.File, err error) {
	var delay lg.Delay
	db := global.GVA_DB.Model(&lg.Delay{})
	if encrypt {
		db = db.Where("elog_encrypt_url = ?", elog)
	} else {
		db = db.Where("elog_url = ?", elog)
	}
	err = db.Preload("ElogFile").Preload("ElogEncryptFile").Order("created_at desc").First(&delay).Error
	if err != nil {
		return lg.File{}, err
	}
	if encrypt {
		file = *delay.ElogEncryptFile
	} else {
		file = *delay.ElogFile
	}
	return
}

func openLetterSync(order lg.Order) {
	var letter lg.Letter
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) (err error) {
		var templateFile lg.File
		if err = tx.Model(&lg.File{}).Where("id = ?", *order.Project.Template.TemplateFileID).First(&templateFile).Error; err != nil {
			return err
		}
		if err = tx.Where("order_no = ?", order.OrderNo).Delete(&lg.Letter{}).Error; err != nil {
			return err
		}
		var file lg.File
		var encryptFile lg.File
		if letter, file, encryptFile, err = lg2.OpenLetter(order, templateFile); err != nil {
			return err
		}
		if err = tx.Create(&file).Error; err != nil {
			return err
		}
		if err = tx.Create(&encryptFile).Error; err != nil {
			return err
		}
		letter.ElogFileID = &file.ID
		letter.ElogEncryptFileID = &encryptFile.ID
		if err = tx.Create(&letter).Error; err != nil {
			return err
		}
		order.LetterID = &letter.ID
		if err = tx.Save(&order).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		global.GVA_LOG.Error("自动化开函失败!", zap.Error(err))
	} else {
		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/lgResultPush"
			var lgResultPush = jrclientrequest.LgResultPush{
				OrderNo:             *letter.OrderNo,
				ElogNo:              *letter.ElogNo,
				InsuranceName:       *letter.InsuranceName,
				InsuranceCreditCode: *letter.InsuranceCreditCode,
				ElogOutDate:         *letter.ElogOutDate,
				ElogUrl:             global.GVA_CONFIG.Insurance.APIDomain + "/letterFileDownload?elog=" + *letter.ElogUrl,
				ElogEncryptUrl:      global.GVA_CONFIG.Insurance.APIDomain + "/letterFileDownload?elog=" + *letter.ElogEncryptUrl + "&type=encrypt",
				TenderDeposit:       *letter.TenderDeposit,
				InsureStartDate:     *letter.InsureStartDate,
				InsureEndDate:       *letter.InsureEndDate,
				InsureDay:           *letter.InsureDay,
				ValidateCode:        *letter.ValidateCode,
			}
			req, err := lg2.GenJRRequest(lgResultPush)
			if err != nil {
				fmt.Println(err)
			}
			var res jrresponse.JRResponse
			client := resty.New()
			for jrResponse := false; !jrResponse; {
				resp, err := client.R().
					SetBody(&req).
					SetResult(&res).
					Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
				if err != nil {
					fmt.Println(err)
				}
				if resp.StatusCode() == http.StatusOK {
					if res.Code != 0 {
						err := errors.New(res.Msg)
						global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
						fmt.Println(err)
					} else {
						byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
						if err != nil {
							fmt.Println(err)
						}
						jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
						if err != nil {
							fmt.Println(err)
						}
						var resData jrclientresponse.Response
						err = json.Unmarshal([]byte(jsonData), &resData)
						if err != nil {
							fmt.Println(err)
						}
						if resData.ReceiveResult != "success" {
							global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
							fmt.Println(errors.New("接收结果不为success"))
						} else {
							jrResponse = true
						}
					}
				}
			}
		}
	}
}
