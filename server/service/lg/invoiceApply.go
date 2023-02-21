package lg

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrclientrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrclientresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrresponse"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	lg2 "github.com/flipped-aurora/gin-vue-admin/server/utils/lg"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"time"
)

type InvoiceApplyService struct {
}

func (invoiceApplyService *InvoiceApplyService) CreateInvoiceApply(invoiceApply lg.InvoiceApply) (err error) {
	err = global.GVA_DB.Create(&invoiceApply).Error
	return err
}

func (invoiceApplyService *InvoiceApplyService) DeleteInvoiceApply(invoiceApply lg.InvoiceApply) (err error) {
	err = global.GVA_DB.Delete(&invoiceApply).Error
	return err
}

func (invoiceApplyService *InvoiceApplyService) DeleteInvoiceApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.InvoiceApply{}, "id in ?", ids.Ids).Error
	return err
}

func (invoiceApplyService *InvoiceApplyService) UpdateInvoiceApply(invoiceApply lg.InvoiceApply) (err error) {
	err = global.GVA_DB.Save(&invoiceApply).Error
	return err
}

func (invoiceApplyService *InvoiceApplyService) GetInvoiceApply(id uint) (invoiceApply lg.InvoiceApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&invoiceApply).Error
	return
}

func (invoiceApplyService *InvoiceApplyService) GetInvoiceApplyInfoList(info lgReq.InvoiceApplySearch) (list []lg.InvoiceApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.InvoiceApply{})
	var invoiceApplys []lg.InvoiceApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Order("created_at desc").Offset(offset).Find(&invoiceApplys).Error
	return invoiceApplys, total, err
}

func (invoiceApplyService *InvoiceApplyService) ApproveInvoiceApply(invoiceApply lg.InvoiceApply) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(2)
		auditOpinion := "受理成功"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		invoiceApply.AuditStatus = &auditStatus
		invoiceApply.AuditOpinion = &auditOpinion
		invoiceApply.AuditDate = &auditDate
		var orderList []jrrequest.Order
		var invoiceList []jrclientrequest.Invoice
		var totalAmount float64
		err = json.Unmarshal([]byte(*invoiceApply.OrderList), &orderList)
		if err != nil {
			return err
		}
		if len(orderList) > 0 {
			totalAmount = 0
			for _, reqOrder := range orderList {
				var order lg.Order
				err = global.GVA_DB.Model(lg.Order{}).Where("order_no = ?", *reqOrder.OrderNo).Preload(clause.Associations).First(&order).Error
				if err != nil {
					return err
				}
				totalAmount += *order.Pay.PayAmount
				var invoiceOrderList []jrclientrequest.Order
				err = json.Unmarshal([]byte(*order.Invoice.OrderList), &invoiceOrderList)
				if err != nil {
					return err
				}
				invoice := jrclientrequest.Invoice{
					InvoiceNo:          *order.Invoice.InvoiceNo,
					InvoiceAmount:      *order.Invoice.InvoiceAmount,
					InvoiceType:        *order.Invoice.InvoiceType,
					InvoiceForm:        *order.Invoice.InvoiceForm,
					InvoicePoint:       *order.Invoice.InvoicePoint,
					InvoiceContent:     *order.Invoice.InvoiceContent,
					InvoiceRemark:      *order.Invoice.InvoiceRemark,
					InvoiceTime:        *order.Invoice.InvoiceTime,
					InvoiceDownloadUrl: *order.Invoice.InvoiceDownloadUrl,
					OrderList:          invoiceOrderList,
				}
				invoiceList = append(invoiceList, invoice)
			}
		}
		jsonInvoiceList, err := json.Marshal(invoiceList)
		if err != nil {
			return err
		}
		invoiceListString := string(jsonInvoiceList)
		invoiceApply.InvoiceList = &invoiceListString
		err = tx.Save(&invoiceApply).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/invoicePush"
			var invoicePush = jrclientrequest.InvoicePush{
				ApplyNo:            *invoiceApply.ApplyNo,
				AuditStatus:        *invoiceApply.AuditStatus,
				AuditOpinion:       *invoiceApply.AuditOpinion,
				AuditDate:          *invoiceApply.AuditDate,
				InvoiceTotalAmount: totalAmount,
				InvoiceList:        invoiceList,
			}
			req, err := lg2.GenJRRequest(invoicePush)
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

func (invoiceApplyService *InvoiceApplyService) RejectInvoiceApply(invoiceApply lg.InvoiceApply) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(3)
		auditOpinion := "受理失败"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		invoiceApply.AuditStatus = &auditStatus
		invoiceApply.AuditOpinion = &auditOpinion
		invoiceApply.AuditDate = &auditDate
		err := tx.Save(&invoiceApply).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomain != "" {
			apiPath := "/jrapi/lg/invoicePush"
			var invoicePush = jrclientrequest.InvoicePush{
				ApplyNo:      *invoiceApply.ApplyNo,
				AuditStatus:  *invoiceApply.AuditStatus,
				AuditOpinion: *invoiceApply.AuditOpinion,
				AuditDate:    *invoiceApply.AuditDate,
			}
			req, err := lg2.GenJRRequest(invoicePush)
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
