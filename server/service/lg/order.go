package lg

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrclientrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrclientresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/nn/nnrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/nn/nnresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/nonmigrate"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	lg2 "github.com/flipped-aurora/gin-vue-admin/server/utils/lg"
	"github.com/go-resty/resty/v2"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type OrderService struct {
}

func (orderService *OrderService) CreateOrder(order lg.Order) (err error) {
	err = global.GVA_DB.Create(&order).Error
	return err
}

func (orderService *OrderService) DeleteOrder(order lg.Order) (err error) {
	err = global.GVA_DB.Delete(&order).Error
	return err
}

func (orderService *OrderService) DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Order{}, "id in ?", ids.Ids).Error
	return err
}

func (orderService *OrderService) UpdateOrder(order lg.Order) (err error) {
	err = global.GVA_DB.Save(&order).Error
	return err
}

func (orderService *OrderService) GetOrder(id uint) (order lg.Order, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&order).Error
	return
}

func (orderService *OrderService) GetOrderInfoList(info lgReq.OrderSearch) (list []lg.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Order{})
	db.Joins("left join lg_apply on lg_apply.id = lg_order.apply_id").
		Joins("left join lg_claim on lg_claim.id = lg_order.claim_id").
		Joins("left join lg_delay on lg_delay.id = lg_order.delay_id").
		Joins("left join lg_letter on lg_letter.id = lg_order.letter_id").
		Joins("left join lg_project on lg_project.id = lg_order.project_id").
		Joins("left join lg_refund on lg_refund.id = lg_order.refund_id")
	var orders []lg.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderNo != nil && *info.OrderNo != "" {
		db = db.Where("lg_order.order_no = ?", info.OrderNo)
	}
	if info.ProjectNo != nil && *info.ProjectNo != "" {
		db = db.Where("lg_apply.project_no = ?", info.ProjectNo)
	}
	if info.ProjectName != nil && *info.ProjectName != "" {
		db = db.Where("lg_apply.project_name = ?", info.ProjectName)
	}
	if info.InsureName != nil && *info.InsureName != "" {
		db = db.Where("lg_apply.insure_name = ?", info.InsureName)
	}
	if info.ElogTemplateId != nil && *info.ElogTemplateId != 0 {
		db = db.Where("lg_project.template_id = ?", info.ElogTemplateId)
	}
	if info.ElogNo != nil && *info.ElogNo != "" {
		db = db.Where("lg_letter.elog_no = ?", info.ElogNo)
	}
	if info.OrderStatus != nil && *info.OrderStatus != "" {
		if *info.OrderStatus == "已撤" {
			db = db.Where("lg_order.revoke_id is not null")
		}
		if *info.OrderStatus == "销函" {
			db = db.Where("lg_order.logout_id is not null")
		}
		if *info.OrderStatus == "理赔" {
			db = db.Where("lg_order.logout_id is null")
			db = db.Where("lg_order.claim_id is not null AND lg_claim.audit_status = 2")
		}
		if *info.OrderStatus == "退函" {
			db = db.Where("lg_order.logout_id is null")
			db = db.Where("lg_order.claim_id is null")
			db = db.Where("lg_order.refund_id is not null AND lg_refund.audit_status = 2")
		}
		if *info.OrderStatus == "延期" {
			db = db.Where("lg_order.logout_id is null")
			db = db.Where("lg_order.claim_id is null")
			db = db.Where("lg_order.refund_id is null")
			db = db.Where("lg_order.delay_id is not null AND lg_delay.audit_status = 2")
		}
		if *info.OrderStatus == "已开" {
			db = db.Where("lg_order.logout_id is null")
			db = db.Where("lg_order.claim_id is null")
			db = db.Where("lg_order.refund_id is null")
			db = db.Where("lg_order.delay_id is null")
			db = db.Where("lg_order.letter_id is not null")
		}
		if *info.OrderStatus == "未开" {
			db = db.Where("lg_order.logout_id is null")
			db = db.Where("lg_order.claim_id is null")
			db = db.Where("lg_order.refund_id is null")
			db = db.Where("lg_order.delay_id is null")
			db = db.Where("lg_order.letter_id is null")
		}
	}
	if info.AuditStatus != nil && *info.AuditStatus != 0 {
		fmt.Println(*info.AuditStatus)
		db = db.Where("lg_apply.audit_status = ?", info.AuditStatus)
	}
	if info.OpenBeginDate != nil {
		db = db.Where("lg_apply.open_begin_date BETWEEN ? AND ?", info.OpenBeginDate[0], info.OpenBeginDate[1])
	}
	if info.ApplyCreatedAt != nil {
		db = db.Where("lg_apply.created_at BETWEEN ? AND ?", info.ApplyCreatedAt[0], info.ApplyCreatedAt[1])
	}
	if info.LetterCreatedAt != nil {
		db = db.Where("lg_letter.created_at BETWEEN ? AND ?", info.LetterCreatedAt[0], info.LetterCreatedAt[1])
	}
	if info.InsureDay != nil && *info.InsureDay != 0 {
		db = db.Where("lg_letter.insure_day = ?", info.InsureDay)
	}
	if info.AuthCode != nil && *info.AuthCode != "" {
		db = db.Where("lg_apply.applicant_auth_code = ?", info.AuthCode)
	}
	if info.AuditDelay != nil {
		db = db.Where("lg_order.delay_id is not null")
	}
	if info.AuditRefund != nil {
		db = db.Where("lg_order.refund_id is not null")
	}
	if info.AuditClaim != nil {
		db = db.Where("lg_order.claim_id is not null")
	}
	if info.IsPayed != nil {
		db = db.Where("lg_order.pay_id is not null")
	}
	if info.EmployeeNo != nil {
		db = db.Where("lg_order.employee_id = ?", info.EmployeeNo)
	}
	if info.NoRevoke != nil {
		db = db.Where("lg_order.revoke_id is null")
	}
	if info.OnlyInvoice != nil {
		db = db.Where("lg_order.invoice_id is not null")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).
		Preload(clause.Associations).
		Preload("Project.Template").
		Preload("Letter.ElogFile").
		Preload("Letter.ElogEncryptFile").
		Preload("Delay.ElogFile").
		Preload("Delay.ElogEncryptFile").
		Order("lg_order.created_at desc").Offset(offset).Find(&orders).Error
	return orders, total, err
}

func (orderService *OrderService) ApproveApply(order lg.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(2)
		auditOpinion := "受理成功"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Apply.AuditStatus = &auditStatus
		order.Apply.AuditOpinion = &auditOpinion
		order.Apply.AuditDate = &auditDate
		realAmount := math.Trunc(*order.Project.TenderDeposit*global.GVA_CONFIG.Insurance.ElogRate*1e2+0.5) * 1e-2
		if realAmount < global.GVA_CONFIG.Insurance.ElogMinAmount {
			order.Apply.RealElogAmount = &global.GVA_CONFIG.Insurance.ElogMinAmount
		} else {
			order.Apply.RealElogAmount = &realAmount
		}
		order.Apply.RealElogRate = &global.GVA_CONFIG.Insurance.ElogRate
		order.Apply.InsuranceName = &global.GVA_CONFIG.Insurance.Name
		order.Apply.InsuranceCreditCode = &global.GVA_CONFIG.Insurance.CreditCode
		err := tx.Save(&order.Apply).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/applyPush"
			var applyPush = jrclientrequest.ApplyPush{
				OrderNo:             *order.OrderNo,
				ApplyNo:             *order.Apply.ApplyNo,
				AuditStatus:         *order.Apply.AuditStatus,
				AuditOpinion:        *order.Apply.AuditOpinion,
				AuditDate:           *order.Apply.AuditDate,
				RealElogAmount:      *order.Apply.RealElogAmount,
				RealElogRate:        *order.Apply.RealElogRate,
				TenderDeposit:       *order.Apply.TenderDeposit,
				InsuranceName:       *order.Apply.InsuranceName,
				InsuranceCreditCode: *order.Apply.InsuredCreditCode,
			}
			req, err := lg2.GenJRRequest(applyPush)
			if err != nil {
				return err
			}
			var res jrresponse.JRResponse
			client := resty.New()
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return err
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return err
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return err
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return err
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return err
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return errors.New("接收结果不为success")
					}
				}
			} else {
				return errors.New("交易中心响应失败")
			}
		}

		return nil
	})
	return err
}

func (orderService *OrderService) RejectApply(order lg.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(3)
		auditOpinion := "受理失败"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Apply.AuditStatus = &auditStatus
		order.Apply.AuditOpinion = &auditOpinion
		order.Apply.AuditDate = &auditDate
		var realAmount float64
		if order.ProjectID != nil {
			realAmount = math.Trunc(*order.Project.TenderDeposit*global.GVA_CONFIG.Insurance.ElogRate*1e2+0.5) * 1e-2
		} else {
			realAmount = math.Trunc(*order.Apply.TenderDeposit*global.GVA_CONFIG.Insurance.ElogRate*1e2+0.5) * 1e-2
		}
		if realAmount < global.GVA_CONFIG.Insurance.ElogMinAmount {
			order.Apply.RealElogAmount = &global.GVA_CONFIG.Insurance.ElogMinAmount
		} else {
			order.Apply.RealElogAmount = &realAmount
		}
		order.Apply.RealElogRate = &global.GVA_CONFIG.Insurance.ElogRate
		order.Apply.InsuranceName = &global.GVA_CONFIG.Insurance.Name
		order.Apply.InsuranceCreditCode = &global.GVA_CONFIG.Insurance.CreditCode
		err := tx.Save(&order.Apply).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/applyPush"
			var applyPush = jrclientrequest.ApplyPush{
				OrderNo:             *order.OrderNo,
				ApplyNo:             *order.Apply.ApplyNo,
				AuditStatus:         *order.Apply.AuditStatus,
				AuditOpinion:        *order.Apply.AuditOpinion,
				AuditDate:           *order.Apply.AuditDate,
				RealElogAmount:      *order.Apply.RealElogAmount,
				RealElogRate:        *order.Apply.RealElogRate,
				TenderDeposit:       *order.Apply.TenderDeposit,
				InsuranceName:       *order.Apply.InsuranceName,
				InsuranceCreditCode: *order.Apply.InsuredCreditCode,
			}
			req, err := lg2.GenJRRequest(applyPush)
			if err != nil {
				return err
			}
			var res jrresponse.JRResponse
			client := resty.New()
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return err
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return err
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return err
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return err
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return err
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return errors.New("接收结果不为success")
					}
				}
			} else {
				return errors.New("交易中心响应失败")
			}
		}

		return nil
	})
	return err
}

func (orderService *OrderService) ApproveDelay(order lg.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(2)
		auditOpinion := "受理成功"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Delay.AuditStatus = &auditStatus
		order.Delay.AuditOpinion = &auditOpinion
		order.Delay.AuditDate = &auditDate

		var templateFile lg.File
		if err = tx.Model(&lg.File{}).Where("id = ?", *order.Project.Template.TemplateFileID).First(&templateFile).Error; err != nil {
			return err
		}

		var letter lg.Letter
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
		order.Delay.ElogFileID = &file.ID
		order.Delay.ElogEncryptFileID = &encryptFile.ID

		order.Delay.ElogUrl = letter.ElogUrl
		order.Delay.ElogEncryptUrl = letter.ElogEncryptUrl
		order.Delay.TenderDeposit = letter.TenderDeposit
		order.Delay.InsureStartDate = letter.InsureStartDate
		order.Delay.InsureEndDate = letter.InsureEndDate
		order.Delay.InsureDay = letter.InsureDay
		order.Delay.ValidateCode = letter.ValidateCode
		err := tx.Save(&order.Delay).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/delayPush"
			var delayPush = jrclientrequest.DelayPush{
				OrderNo:         *order.OrderNo,
				ApplyNo:         *order.Delay.ApplyNo,
				ElogNo:          *order.Delay.ElogNo,
				AuditStatus:     *order.Delay.AuditStatus,
				AuditOpinion:    *order.Delay.AuditOpinion,
				AuditDate:       *order.Delay.AuditDate,
				ElogUrl:         global.GVA_CONFIG.Insurance.APIDomain + "/delayFileDownload?elog=" + *letter.ElogUrl,
				ElogEncryptUrl:  global.GVA_CONFIG.Insurance.APIDomain + "/delayFileDownload?elog=" + *letter.ElogEncryptUrl + "&type=encrypt",
				TenderDeposit:   *letter.TenderDeposit,
				InsureStartDate: *letter.InsureStartDate,
				InsureEndDate:   *letter.InsureEndDate,
				InsureDay:       *letter.InsureDay,
				ValidateCode:    *letter.ValidateCode,
			}
			req, err := lg2.GenJRRequest(delayPush)
			if err != nil {
				return err
			}
			var res jrresponse.JRResponse
			client := resty.New()
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return err
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return err
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return err
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return err
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return err
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return errors.New("接收结果不为success")
					}
				}
			} else {
				return errors.New("交易中心响应失败")
			}
		}

		return nil
	})
	return err
}

func (orderService *OrderService) RejectDelay(order lg.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(3)
		auditOpinion := "受理失败"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Delay.AuditStatus = &auditStatus
		order.Delay.AuditOpinion = &auditOpinion
		order.Delay.AuditDate = &auditDate

		var templateFile lg.File
		if err = tx.Model(&lg.File{}).Where("id = ?", *order.Project.Template.TemplateFileID).First(&templateFile).Error; err != nil {
			return err
		}
		err := tx.Save(&order.Delay).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/delayPush"
			var delayPush = jrclientrequest.DelayPush{
				OrderNo:         *order.OrderNo,
				ApplyNo:         *order.Delay.ApplyNo,
				ElogNo:          *order.Letter.ElogNo,
				AuditStatus:     *order.Delay.AuditStatus,
				AuditOpinion:    *order.Delay.AuditOpinion,
				AuditDate:       *order.Delay.AuditDate,
				ElogUrl:         *order.Letter.ElogUrl,
				ElogEncryptUrl:  *order.Letter.ElogEncryptUrl,
				TenderDeposit:   *order.Letter.TenderDeposit,
				InsureStartDate: *order.Letter.InsureStartDate,
				InsureEndDate:   *order.Letter.InsureEndDate,
				InsureDay:       *order.Letter.InsureDay,
				ValidateCode:    *order.Letter.ValidateCode,
			}
			req, err := lg2.GenJRRequest(delayPush)
			if err != nil {
				return err
			}
			var res jrresponse.JRResponse
			client := resty.New()
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return err
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return err
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return err
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return err
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return err
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return errors.New("接收结果不为success")
					}
				}
			} else {
				return errors.New("交易中心响应失败")
			}
		}

		return nil
	})
	return err
}

func (orderService *OrderService) ApproveRefund(order lg.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(2)
		auditOpinion := "受理成功"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Refund.AuditStatus = &auditStatus
		order.Refund.AuditOpinion = &auditOpinion
		order.Refund.AuditDate = &auditDate
		order.Refund.PayAmount = order.Pay.PayAmount
		err := tx.Save(&order.Refund).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/refundPush"
			var refundPush = jrclientrequest.RefundPush{
				OrderNo:      *order.OrderNo,
				ApplyNo:      *order.Refund.ApplyNo,
				ElogNo:       *order.Refund.ElogNo,
				AuditStatus:  *order.Refund.AuditStatus,
				AuditOpinion: *order.Refund.AuditOpinion,
				AuditDate:    *order.Refund.AuditDate,
				PayAmount:    *order.Refund.PayAmount,
			}
			req, err := lg2.GenJRRequest(refundPush)
			if err != nil {
				return err
			}
			var res jrresponse.JRResponse
			client := resty.New()
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return err
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return err
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return err
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return err
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return err
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return errors.New("接收结果不为success")
					}
				}
			} else {
				return errors.New("交易中心响应失败")
			}
		}

		return nil
	})
	return err
}

func (orderService *OrderService) RejectRefund(order lg.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(3)
		auditOpinion := "受理失败"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Refund.AuditStatus = &auditStatus
		order.Refund.AuditOpinion = &auditOpinion
		order.Refund.AuditDate = &auditDate
		order.Refund.PayAmount = order.Pay.PayAmount
		err := tx.Save(&order.Refund).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/refundPush"
			var refundPush = jrclientrequest.RefundPush{
				OrderNo:      *order.OrderNo,
				ApplyNo:      *order.Refund.ApplyNo,
				ElogNo:       *order.Refund.ElogNo,
				AuditStatus:  *order.Refund.AuditStatus,
				AuditOpinion: *order.Refund.AuditOpinion,
				AuditDate:    *order.Refund.AuditDate,
			}
			req, err := lg2.GenJRRequest(refundPush)
			if err != nil {
				return err
			}
			var res jrresponse.JRResponse
			client := resty.New()
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return err
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return err
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return err
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return err
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return err
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return errors.New("接收结果不为success")
					}
				}
			} else {
				return errors.New("交易中心响应失败")
			}
		}

		return nil
	})
	return err
}

func (orderService *OrderService) ApproveClaim(order lg.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(2)
		auditOpinion := "受理成功"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Claim.AuditStatus = &auditStatus
		order.Claim.AuditOpinion = &auditOpinion
		order.Claim.AuditDate = &auditDate
		order.Claim.RealClaimAmount = order.Claim.ClaimAmount
		order.Claim.RealClaimDate = &auditDate
		err := tx.Save(&order.Claim).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/claimPush"
			var claimPush = jrclientrequest.ClaimPush{
				OrderNo:         *order.OrderNo,
				ApplyNo:         *order.Claim.ApplyNo,
				ElogNo:          *order.Claim.ElogNo,
				AuditStatus:     *order.Claim.AuditStatus,
				AuditOpinion:    *order.Claim.AuditOpinion,
				AuditDate:       *order.Claim.AuditDate,
				RealClaimAmount: *order.Claim.RealClaimAmount,
				RealClaimDate:   *order.Claim.RealClaimDate,
			}
			req, err := lg2.GenJRRequest(claimPush)
			if err != nil {
				return err
			}
			var res jrresponse.JRResponse
			client := resty.New()
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return err
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return err
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return err
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return err
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return err
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return errors.New("接收结果不为success")
					}
				}
			} else {
				return errors.New("交易中心响应失败")
			}
		}

		return nil
	})
	return err
}

func (orderService *OrderService) RejectClaim(order lg.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(3)
		auditOpinion := "受理失败"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Claim.AuditStatus = &auditStatus
		order.Claim.AuditOpinion = &auditOpinion
		order.Claim.AuditDate = &auditDate
		err := tx.Save(&order.Claim).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/claimPush"
			var claimPush = jrclientrequest.ClaimPush{
				OrderNo:      *order.OrderNo,
				ApplyNo:      *order.Claim.ApplyNo,
				ElogNo:       *order.Claim.ElogNo,
				AuditStatus:  *order.Claim.AuditStatus,
				AuditOpinion: *order.Claim.AuditOpinion,
				AuditDate:    *order.Claim.AuditDate,
			}
			req, err := lg2.GenJRRequest(claimPush)
			if err != nil {
				return err
			}
			var res jrresponse.JRResponse
			client := resty.New()
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return err
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return err
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return err
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return err
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return err
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return errors.New("接收结果不为success")
					}
				}
			} else {
				return errors.New("交易中心响应失败")
			}
		}

		return nil
	})
	return err
}

func (orderService *OrderService) OpenLetter(order lg.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var templateFile lg.File
		if err = tx.Model(&lg.File{}).Where("id = ?", *order.Project.Template.TemplateFileID).First(&templateFile).Error; err != nil {
			return err
		}
		if err = tx.Where("order_no = ?", order.OrderNo).Delete(&lg.Letter{}).Error; err != nil {
			return err
		}
		var letter lg.Letter
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
				return err
			}
			var res jrresponse.JRResponse
			client := resty.New()
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return err
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return err
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return err
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return err
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return err
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return errors.New("接收结果不为success")
					}
				}
			} else {
				return errors.New("交易中心响应失败")
			}
		}

		return nil
	})
	return err
}

func (orderService *OrderService) RePush(order lg.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var templateFile lg.File
		if err = tx.Model(&lg.File{}).Where("id = ?", *order.Project.Template.TemplateFileID).First(&templateFile).Error; err != nil {
			return err
		}
		if err = tx.Where("order_no = ?", order.OrderNo).Delete(&lg.Letter{}).Error; err != nil {
			return err
		}
		var letter lg.Letter
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
				return err
			}
			var res jrresponse.JRResponse
			client := resty.New()
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return err
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return err
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return err
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return err
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return err
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return errors.New("接收结果不为success")
					}
				}
			} else {
				return errors.New("交易中心响应失败")
			}
		}
		isRepushed := true
		order.IsRepushed = &isRepushed
		if err = tx.Save(&order).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

func (orderService *OrderService) GetOrderStatisticData() (orderStatisticData nonmigrate.OrderStatisticData, err error) {
	db := global.GVA_DB.Model(&lg.Order{})
	db.Joins("Pay").Joins("Letter").Joins("Refund").Joins("Claim")
	//未撤函，未理赔，未退函
	db = db.Where("lg_order.revoke_id is null")
	db = db.Where("lg_order.claim_id is null")
	db = db.Where("lg_order.refund_id is null")

	err = db.Select("COALESCE(SUM(Letter.tender_deposit), 0) as TotalGuaranteeAmount, COALESCE(SUM(Pay.pay_amount), 0) as TotalElogAmount").Scan(&orderStatisticData).Error

	return orderStatisticData, err
}

func (orderService *OrderService) ExportExcel(info lgReq.OrderSearch) (excelData []byte, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Order{})
	db.Joins("left join lg_apply on lg_apply.id = lg_order.apply_id").
		Joins("left join lg_claim on lg_claim.id = lg_order.claim_id").
		Joins("left join lg_delay on lg_delay.id = lg_order.delay_id").
		Joins("left join lg_letter on lg_letter.id = lg_order.letter_id").
		Joins("left join lg_project on lg_project.id = lg_order.project_id").
		Joins("left join lg_refund on lg_refund.id = lg_order.refund_id")
	var orders []lg.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderNo != nil && *info.OrderNo != "" {
		db = db.Where("lg_order.order_no = ?", info.OrderNo)
	}
	if info.ProjectNo != nil && *info.ProjectNo != "" {
		db = db.Where("lg_apply.project_no = ?", info.ProjectNo)
	}
	if info.ProjectName != nil && *info.ProjectName != "" {
		db = db.Where("lg_apply.project_name = ?", info.ProjectName)
	}
	if info.InsureName != nil && *info.InsureName != "" {
		db = db.Where("lg_apply.insure_name = ?", info.InsureName)
	}
	if info.ElogTemplateId != nil && *info.ElogTemplateId != 0 {
		db = db.Where("lg_project.template_id = ?", info.ElogTemplateId)
	}
	if info.ElogNo != nil && *info.ElogNo != "" {
		db = db.Where("lg_letter.elog_no = ?", info.ElogNo)
	}
	if info.OrderStatus != nil && *info.OrderStatus != "" {
		if *info.OrderStatus == "已撤" {
			db = db.Where("lg_order.revoke_id is not null")
		}
		if *info.OrderStatus == "销函" {
			db = db.Where("lg_order.logout_id is not null")
		}
		if *info.OrderStatus == "理赔" {
			db = db.Where("lg_order.logout_id is null")
			db = db.Where("lg_order.claim_id is not null AND lg_claim.audit_status = 2")
		}
		if *info.OrderStatus == "退函" {
			db = db.Where("lg_order.logout_id is null")
			db = db.Where("lg_order.claim_id is null")
			db = db.Where("lg_order.refund_id is not null AND lg_refund.audit_status = 2")
		}
		if *info.OrderStatus == "延期" {
			db = db.Where("lg_order.logout_id is null")
			db = db.Where("lg_order.claim_id is null")
			db = db.Where("lg_order.refund_id is null")
			db = db.Where("lg_order.delay_id is not null AND lg_delay.audit_status = 2")
		}
		if *info.OrderStatus == "已开" {
			db = db.Where("lg_order.logout_id is null")
			db = db.Where("lg_order.claim_id is null")
			db = db.Where("lg_order.refund_id is null")
			db = db.Where("lg_order.delay_id is null")
			db = db.Where("lg_order.letter_id is not null")
		}
		if *info.OrderStatus == "未开" {
			db = db.Where("lg_order.logout_id is null")
			db = db.Where("lg_order.claim_id is null")
			db = db.Where("lg_order.refund_id is null")
			db = db.Where("lg_order.delay_id is null")
			db = db.Where("lg_order.letter_id is null")
		}
	}
	if info.AuditStatus != nil && *info.AuditStatus != 0 {
		db = db.Where("lg_apply.audit_status = ?", info.AuditStatus)
	}
	if info.OpenBeginDate != nil {
		db = db.Where("lg_apply.open_begin_date BETWEEN ? AND ?", info.OpenBeginDate[0], info.OpenBeginDate[1])
	}
	if info.ApplyCreatedAt != nil {
		db = db.Where("lg_apply.created_at BETWEEN ? AND ?", info.ApplyCreatedAt[0], info.ApplyCreatedAt[1])
	}
	if info.LetterCreatedAt != nil {
		db = db.Where("lg_letter.created_at BETWEEN ? AND ?", info.LetterCreatedAt[0], info.LetterCreatedAt[1])
	}
	if info.InsureDay != nil && *info.InsureDay != 0 {
		db = db.Where("lg_letter.insure_day = ?", info.InsureDay)
	}
	if info.AuthCode != nil && *info.AuthCode != "" {
		db = db.Where("lg_apply.applicant_auth_code = ?", info.AuthCode)
	}
	if info.AuditDelay != nil {
		db = db.Where("lg_order.delay_id is not null")
	}
	if info.AuditRefund != nil {
		db = db.Where("lg_order.refund_id is not null")
	}
	if info.AuditClaim != nil {
		db = db.Where("lg_order.claim_id is not null")
	}
	if info.IsPayed != nil {
		db = db.Where("lg_order.pay_id is not null")
	}
	if info.EmployeeNo != nil {
		db = db.Where("lg_order.employee_id = ?", info.EmployeeNo)
	}
	if info.NoRevoke != nil {
		db = db.Where("lg_order.revoke_id is null")
	}
	if info.OnlyInvoice != nil {
		db = db.Where("lg_order.invoice_id is not null")
	}
	err = db.Limit(limit).Preload(clause.Associations).Order("lg_order.created_at desc").Offset(offset).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	excel := excelize.NewFile()
	_ = excel.SetSheetRow("Sheet1", "A1", &[]string{"保函文件下载", "交易中心", "保函申请编码", "申请企业", "标段名称", "标段编号", "受益人名称", "担保金额（元）", "保函起始日期", "保函截止日期", "订单状态", "开标时间", "保费金额", "所属市", "所属县", "审核时间", "申请日期", "开函日期", "保函编码", "审核状态", "付款时间", "付款金额", "交易单号", "工号"})
	for i, order := range orders {
		axis := fmt.Sprintf("A%d", i+2)
		elogAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", *order.Apply.TenderDeposit**order.Apply.ProductRate), 64)
		elogUrl := ""
		insureStartDate := ""
		insureEndDate := ""
		letterOpenDate := ""
		elogNo := ""
		if order.LetterID != nil {
			elogUrl = global.GVA_CONFIG.Insurance.APIDomain + "letterFileDownload?elog=" + *order.Letter.ElogUrl
			insureStartDate = *order.Letter.InsureStartDate
			insureEndDate = *order.Letter.InsureEndDate
			letterOpenDate = order.Letter.CreatedAt.Format("2006-01-02 15:04:05")
			elogNo = *order.Letter.ElogNo
		}
		projectCity := ""
		projectCounty := ""
		if order.ProjectID != nil {
			if order.Project.ProjectCity != nil {
				projectCity = *order.Project.ProjectCity
			}
			if order.Project.ProjectCounty != nil {
				projectCounty = *order.Project.ProjectCounty
			}
		} else {
			projectCity = ""
			projectCounty = ""
		}
		var payTime string
		var payAmount float64
		var payTransNo string
		if order.PayID != nil {
			payTime = *order.Pay.PayTime
			payAmount = *order.Pay.PayAmount
			payTransNo = *order.Pay.PayTransNo
		} else {
			payTime = ""
			payAmount = 0.0
			payTransNo = ""
		}
		var employeeNo string
		if order.EmployeeID != nil {
			employeeNo = *order.Employee.EmployeeNo
		} else {
			employeeNo = ""
		}
		_ = excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			elogUrl,
			"江西云平台",
			*order.Apply.ApplyNo,
			*order.Apply.InsureName,
			*order.Apply.ProjectName,
			*order.Apply.ProjectNo,
			*order.Apply.InsuredName,
			*order.Apply.TenderDeposit,
			insureStartDate,
			insureEndDate,
			lg2.OrderStatus(order),
			*order.Apply.OpenBeginDate,
			elogAmount,
			projectCity,
			projectCounty,
			*order.Apply.AuditDate,
			order.Apply.CreatedAt.Format("2006-01-02 15:04:05"),
			letterOpenDate,
			elogNo,
			lg2.AuditStatus(*order.Apply.AuditStatus),
			payTime,
			payAmount,
			payTransNo,
			employeeNo,
		})
	}
	excelBuffer, err := excel.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	excelData = excelBuffer.Bytes()
	return
}

func (orderService *OrderService) GetOrderByNos(orderNos []string) (orders []lg.Order, err error) {
	err = global.GVA_DB.Where("order_no in ?", orderNos).Preload(clause.Associations).Find(&orders).Error
	return
}

func (orderService *OrderService) RequestInvoice(reqInvoice nnrequest.NNRequestInvoice) (err error) {
	type NNDetail struct {
		GoodsName   string `json:"goodsName"`
		WithTaxFlag string `json:"withTaxFlag"`
		Price       string `json:"price"`
		Num         string `json:"num"`
		TaxRate     string `json:"taxRate"`
	}
	type NNOrder struct {
		BuyerName     string     `json:"buyerName"`
		BuyerTaxNum   *string    `json:"buyerTaxNum,omitempty"`
		BuyerTel      *string    `json:"buyerTel,omitempty"`
		BuyerAddress  *string    `json:"buyerAddress,omitempty"`
		BuyerAccount  *string    `json:"buyerAccount,omitempty"`
		SalerTaxNum   string     `json:"salerTaxNum"`
		SalerTel      string     `json:"salerTel"`
		SalerAddress  string     `json:"salerAddress"`
		SalerAccount  *string    `json:"salerAccount,omitempty"`
		OrderNo       string     `json:"orderNo"`
		InvoiceDate   string     `json:"invoiceDate"`
		Remark        *string    `json:"remark,omitempty"`
		Checker       *string    `json:"checker,omitempty"`
		Payee         *string    `json:"payee,omitempty"`
		Clerk         string     `json:"clerk"`
		PushMode      string     `json:"pushMode"`
		BuyerPhone    *string    `json:"buyerPhone,omitempty"`
		InvoiceType   string     `json:"invoiceType"`
		InvoiceDetail []NNDetail `json:"invoiceDetail"`
	}
	type NNReq struct {
		Order NNOrder `json:"order"`
	}
	bankName := ""
	bankNo := ""
	if reqInvoice.InvoiceApply.BankName != nil {
		bankName = *reqInvoice.InvoiceApply.BankName
	}
	if reqInvoice.InvoiceApply.BankNo != nil {
		bankName = *reqInvoice.InvoiceApply.BankNo
	}
	buyerAccount := bankName + bankNo
	salerAccount := global.GVA_CONFIG.Insurance.BankName + global.GVA_CONFIG.Insurance.BankNo
	checker := global.GVA_CONFIG.Insurance.NNChecker
	payee := global.GVA_CONFIG.Insurance.NNPayee
	timestamp := time.Now().Unix()
	rand.Seed(timestamp)
	randomOrder := rand.Intn(899999) + 100000
	pushMode := -1
	if reqInvoice.InvoiceApply.CompanyTel != nil {
		pushMode = 1
	}
	nnreq := NNReq{
		Order: NNOrder{
			BuyerName:    *reqInvoice.InvoiceApply.InvoiceTile,
			BuyerTaxNum:  reqInvoice.InvoiceApply.TaxNo,
			BuyerTel:     reqInvoice.InvoiceApply.CompanyTel,
			BuyerAddress: reqInvoice.InvoiceApply.CompanyAddress,
			BuyerAccount: &buyerAccount,
			SalerTaxNum:  global.GVA_CONFIG.Insurance.NNTaxNo,
			SalerTel:     global.GVA_CONFIG.Insurance.Tel,
			SalerAddress: global.GVA_CONFIG.Insurance.Address,
			SalerAccount: &salerAccount,
			OrderNo:      time.Now().Format("20060102150405") + strconv.Itoa(randomOrder),
			InvoiceDate:  time.Now().Format("2006-01-02 15:04:05"),
			Remark:       reqInvoice.InvoiceApply.Remarks,
			Checker:      &checker,
			Payee:        &payee,
			Clerk:        global.GVA_CONFIG.Insurance.NNClerk,
			PushMode:     strconv.Itoa(pushMode),
			BuyerPhone:   reqInvoice.InvoiceApply.CompanyTel,
			InvoiceType:  "1",
			InvoiceDetail: []NNDetail{{
				GoodsName:   "电子保函",
				WithTaxFlag: "1",
				Price:       strconv.FormatFloat(*reqInvoice.Order.Pay.PayAmount, 'f', 8, 64),
				Num:         "1",
				TaxRate:     strconv.FormatFloat(global.GVA_CONFIG.Insurance.NNTaxRate, 'f', 2, 64),
			}},
		},
	}
	jsonReq, _ := json.Marshal(&nnreq)
	res, err := lg2.NNSendRequest("nuonuo.ElectronInvoice.requestBillingNew", string(jsonReq))
	if err != nil {
		return
	}
	var resInvoice nnresponse.NNRequestInvoice
	err = json.Unmarshal(res, &resInvoice)
	if err != nil {
		return
	}
	if resInvoice.Code != "E0000" {
		return errors.New(resInvoice.Code + ":" + resInvoice.Describe)
	}
	invoiceForm := "B1"
	invoicePoint, _ := strconv.ParseFloat(fmt.Sprintf("%.6f", global.GVA_CONFIG.Insurance.NNTaxRate*100), 64)
	invoiceContent := "发票内容"
	type Result struct {
		InvoiceSerialNum string `json:"invoiceSerialNum"`
	}
	type ResultRequestInvoice struct {
		Code     string `json:"code"`
		Describe string `json:"describe"`
		Result   Result `json:"result"`
	}
	var resultRequestInvoice ResultRequestInvoice
	err = json.Unmarshal(res, &resultRequestInvoice)
	if err != nil {
		return
	}
	orderList := `[{"orderNo":"` + *reqInvoice.Order.OrderNo + `","orderInvoiceAmount":` + strconv.FormatFloat(*reqInvoice.Order.Pay.PayAmount, 'f', 2, 64) + `}]`
	invoice := lg.Invoice{
		InvoiceNo:          &nnreq.Order.OrderNo,
		InvoiceAmount:      reqInvoice.Order.Pay.PayAmount,
		InvoiceType:        reqInvoice.InvoiceApply.InvoiceType,
		InvoiceTileType:    reqInvoice.InvoiceApply.InvoiceType,
		InvoiceTile:        reqInvoice.InvoiceApply.InvoiceTile,
		TaxNo:              reqInvoice.InvoiceApply.TaxNo,
		BankName:           reqInvoice.InvoiceApply.BankName,
		BankNo:             reqInvoice.InvoiceApply.BankNo,
		CompanyAddress:     reqInvoice.InvoiceApply.CompanyAddress,
		CompanyTel:         reqInvoice.InvoiceApply.CompanyTel,
		Remarks:            reqInvoice.InvoiceApply.Remarks,
		InvoiceForm:        &invoiceForm,
		InvoicePoint:       &invoicePoint,
		InvoiceContent:     &invoiceContent,
		InvoiceRemark:      reqInvoice.InvoiceApply.Remarks,
		InvoiceTime:        &nnreq.Order.InvoiceDate,
		InvoiceSerialNum:   &resultRequestInvoice.Result.InvoiceSerialNum,
		InvoiceDownloadUrl: nil,
		OrderList:          &orderList,
	}
	err = global.GVA_DB.Create(&invoice).Error
	if err != nil {
		return
	}
	return global.GVA_DB.Model(&lg.Order{}).Where("id = ?", reqInvoice.Order.ID).Update("invoice_id", invoice.ID).Error
}

func (orderService *OrderService) QueryInvoice(reqInvoice nnrequest.NNQueryInvoice) (err error) {
	type NNReq struct {
		SerialNos            []string `json:"serialNos,omitempty"`
		OrderNos             []string `json:"orderNos,omitempty"`
		IsOfferInvoiceDetail string   `json:"isOfferInvoiceDetail,omitempty"`
	}
	nnreq := NNReq{
		SerialNos: []string{*reqInvoice.Order.Invoice.InvoiceSerialNum},
	}
	jsonReq, _ := json.Marshal(&nnreq)
	res, err := lg2.NNSendRequest("nuonuo.ElectronInvoice.queryInvoiceResult", string(jsonReq))
	if err != nil {
		return
	}
	var resInvoice nnresponse.NNRequestInvoice
	err = json.Unmarshal(res, &resInvoice)
	if err != nil {
		return
	}
	if resInvoice.Code != "E0000" {
		return errors.New(resInvoice.Code + ":" + resInvoice.Describe)
	}
	type Result struct {
		PdfUrl string `json:"pdfUrl"`
	}
	type ResultRequestInvoice struct {
		Code     string   `json:"code"`
		Describe string   `json:"describe"`
		Result   []Result `json:"result"`
	}
	var resultRequestInvoice ResultRequestInvoice
	err = json.Unmarshal(res, &resultRequestInvoice)
	if err != nil {
		return
	}
	reqInvoice.Order.Invoice.InvoiceDownloadUrl = &resultRequestInvoice.Result[0].PdfUrl
	return global.GVA_DB.Save(&reqInvoice.Order.Invoice).Error
}

func (orderService *OrderService) AssignOrder(assign lgReq.AssignOrder) (err error) {
	var order lg.Order
	err = global.GVA_DB.Where("order_no = ?", assign.OrderNo).First(&order).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // 判断人员编号是否已经存在
		return errors.New("该订单编号不存在")
	} else {
		return err
	}
	if order.EmployeeID != nil {
		return errors.New("该订单已经分配")
	} else {
		return global.GVA_DB.Model(&order).Where("id = ?", order.ID).Update("employee_id", assign.EmployeeId).Error
	}
}
