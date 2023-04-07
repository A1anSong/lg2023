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
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/nn/nnresponse"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	lg2 "github.com/flipped-aurora/gin-vue-admin/server/utils/lg"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
	"net/http"
	"strconv"
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
	if info.InvoiceTile != nil {
		db = db.Where("invoice_tile = ?", info.InvoiceTile)
	}
	if info.ApplyTime != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.ApplyTime[0], info.ApplyTime[1])
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

func AutoMaticAuditInvoiceApply() {
	var invoiceApplyList []lg.InvoiceApply
	err := global.GVA_DB.Where("audit_status is null").Find(&invoiceApplyList).Error
	if err != nil {
		return
	}
	for _, invoiceApply := range invoiceApplyList {
		go AuditInvoiceApply(invoiceApply)
	}
}

func AuditInvoiceApply(invoiceApply lg.InvoiceApply) {
	var orderList []jrrequest.Order
	err := json.Unmarshal([]byte(*invoiceApply.OrderList), &orderList)
	if err != nil {
		return
	}
	var ordersString []string
	var orders []lg.Order
	if len(orderList) > 0 {
		for _, order := range orderList {
			ordersString = append(ordersString, *order.OrderNo)
		}
		err = global.GVA_DB.Where("order_no in ?", ordersString).Preload(clause.Associations).Find(&orders).Error
		if err != nil {
			return
		}
		for _, order := range orders {
			if *order.IsOfflineRefund == true {
				go AuditRejectInvoiceApply(invoiceApply, "订单"+*order.OrderNo+"已经线下退款")
				return
			}
		}
		for _, order := range orders {
			projectOpenTime, _ := time.Parse("2006-01-02 15:04:05", *order.Project.ProjectOpenTime)
			if *order.Project.ProjectOpenTime > time.Now().Format("2006-01-02") {
				if time.Now().Sub(projectOpenTime).Hours() < 72 {
					return
				}
			}
		}
		isAllInvoiceReady := true
		for _, order := range orders {
			if order.LetterID == nil {
				go RequestInvoice(order, invoiceApply)
				isAllInvoiceReady = false
			}
		}
		if isAllInvoiceReady {
			go AuditApproveInvoiceApply(invoiceApply)
		}
	}
}

func AuditRejectInvoiceApply(invoiceApply lg.InvoiceApply, auditOpinion string) {
	auditStatus := int64(3)
	auditDate := time.Now().Format("2006-01-02 15:04:05")
	invoiceApply.AuditStatus = &auditStatus
	invoiceApply.AuditOpinion = &auditOpinion
	invoiceApply.AuditDate = &auditDate
	err := global.GVA_DB.Save(&invoiceApply).Error
	if err != nil {
		return
	} else {
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
				return
			}
			var res jrresponse.JRResponse
			client := resty.New()
			for jrResponse := false; !jrResponse; {
				resp, err := client.R().
					SetBody(&req).
					SetResult(&res).
					Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
				if err != nil {
					return
				}
				if resp.StatusCode() == http.StatusOK {
					if res.Code != 0 {
						err := errors.New(res.Msg)
						global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
						return
					} else {
						byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
						if err != nil {
							return
						}
						jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
						if err != nil {
							return
						}
						var resData jrclientresponse.Response
						err = json.Unmarshal([]byte(jsonData), &resData)
						if err != nil {
							return
						}
						if resData.ReceiveResult != "success" {
							global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
							return
						} else {
							jrResponse = true
						}
					}
				}
			}
		}
	}
}

func RequestInvoice(order lg.Order, invoiceApply lg.InvoiceApply) {
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
	if invoiceApply.BankName != nil {
		bankName = *invoiceApply.BankName
	}
	if invoiceApply.BankNo != nil {
		bankName = *invoiceApply.BankNo
	}
	buyerAccount := bankName + bankNo
	salerAccount := global.GVA_CONFIG.Insurance.BankName + global.GVA_CONFIG.Insurance.BankNo
	checker := global.GVA_CONFIG.Insurance.NNChecker
	payee := global.GVA_CONFIG.Insurance.NNPayee
	timestamp := time.Now().Unix()
	rand.Seed(timestamp)
	randomOrder := rand.Intn(899999) + 100000
	pushMode := -1
	if invoiceApply.CompanyTel != nil {
		pushMode = 1
	}
	nnreq := NNReq{
		Order: NNOrder{
			BuyerName:    *invoiceApply.InvoiceTile,
			BuyerTaxNum:  invoiceApply.TaxNo,
			BuyerTel:     invoiceApply.CompanyTel,
			BuyerAddress: invoiceApply.CompanyAddress,
			BuyerAccount: &buyerAccount,
			SalerTaxNum:  global.GVA_CONFIG.Insurance.NNTaxNo,
			SalerTel:     global.GVA_CONFIG.Insurance.Tel,
			SalerAddress: global.GVA_CONFIG.Insurance.Address,
			SalerAccount: &salerAccount,
			OrderNo:      time.Now().Format("20060102150405") + strconv.Itoa(randomOrder),
			InvoiceDate:  time.Now().Format("2006-01-02 15:04:05"),
			Remark:       invoiceApply.Remarks,
			Checker:      &checker,
			Payee:        &payee,
			Clerk:        global.GVA_CONFIG.Insurance.NNClerk,
			PushMode:     strconv.Itoa(pushMode),
			BuyerPhone:   invoiceApply.CompanyTel,
			InvoiceType:  "1",
			InvoiceDetail: []NNDetail{{
				GoodsName:   "电子保函",
				WithTaxFlag: "1",
				Price:       strconv.FormatFloat(*order.Pay.PayAmount, 'f', 8, 64),
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
		return
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
	orderList := `[{"orderNo":"` + *order.OrderNo + `","orderInvoiceAmount":` + strconv.FormatFloat(*order.Pay.PayAmount, 'f', 2, 64) + `}]`
	invoice := lg.Invoice{
		InvoiceNo:          &nnreq.Order.OrderNo,
		InvoiceAmount:      order.Pay.PayAmount,
		InvoiceType:        invoiceApply.InvoiceType,
		InvoiceTileType:    invoiceApply.InvoiceType,
		InvoiceTile:        invoiceApply.InvoiceTile,
		TaxNo:              invoiceApply.TaxNo,
		BankName:           invoiceApply.BankName,
		BankNo:             invoiceApply.BankNo,
		CompanyAddress:     invoiceApply.CompanyAddress,
		CompanyTel:         invoiceApply.CompanyTel,
		Remarks:            invoiceApply.Remarks,
		InvoiceForm:        &invoiceForm,
		InvoicePoint:       &invoicePoint,
		InvoiceContent:     &invoiceContent,
		InvoiceRemark:      invoiceApply.Remarks,
		InvoiceTime:        &nnreq.Order.InvoiceDate,
		InvoiceSerialNum:   &resultRequestInvoice.Result.InvoiceSerialNum,
		InvoiceDownloadUrl: nil,
		OrderList:          &orderList,
	}
	for isSaveSuccess := false; !isSaveSuccess; {
		err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			err = tx.Create(&invoice).Error
			if err != nil {
				return err
			}
			err = tx.Model(&lg.Order{}).Where("id = ?", order.ID).Update("invoice_id", invoice.ID).Error
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			global.GVA_LOG.Error("创建发票失败", zap.Error(err))
		} else {
			isSaveSuccess = true
		}
	}
}

func AuditApproveInvoiceApply(invoiceApply lg.InvoiceApply) {
	auditStatus := int64(2)
	auditOpinion := "受理成功"
	auditDate := time.Now().Format("2006-01-02 15:04:05")
	invoiceApply.AuditStatus = &auditStatus
	invoiceApply.AuditOpinion = &auditOpinion
	invoiceApply.AuditDate = &auditDate
	var orderList []jrrequest.Order
	var invoiceList []jrclientrequest.Invoice
	var totalAmount float64
	err := json.Unmarshal([]byte(*invoiceApply.OrderList), &orderList)
	if err != nil {
		return
	}
	if len(orderList) > 0 {
		totalAmount = 0
		for _, reqOrder := range orderList {
			var order lg.Order
			err = global.GVA_DB.Model(lg.Order{}).Where("order_no = ?", *reqOrder.OrderNo).Preload(clause.Associations).First(&order).Error
			if err != nil {
				return
			}
			totalAmount += *order.Pay.PayAmount
			var invoiceOrderList []jrclientrequest.Order
			err = json.Unmarshal([]byte(*order.Invoice.OrderList), &invoiceOrderList)
			if err != nil {
				return
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
		return
	}
	invoiceListString := string(jsonInvoiceList)
	invoiceApply.InvoiceList = &invoiceListString
	err = global.GVA_DB.Save(&invoiceApply).Error
	if err != nil {
		return
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
			return
		}
		var res jrresponse.JRResponse
		client := resty.New()
		for jrResponse := false; !jrResponse; {
			resp, err := client.R().
				SetBody(&req).
				SetResult(&res).
				Post(global.GVA_CONFIG.Insurance.JRAPIDomain + apiPath)
			if err != nil {
				return
			}
			if resp.StatusCode() == http.StatusOK {
				if res.Code != 0 {
					err := errors.New(res.Msg)
					global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
					return
				} else {
					byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
					if err != nil {
						return
					}
					jsonData, err := lg2.Sm4Decrypt(byteEncryptData)
					if err != nil {
						return
					}
					var resData jrclientresponse.Response
					err = json.Unmarshal([]byte(jsonData), &resData)
					if err != nil {
						return
					}
					if resData.ReceiveResult != "success" {
						global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
						return
					} else {
						jrResponse = true
					}
				}
			}
		}
	}
}
