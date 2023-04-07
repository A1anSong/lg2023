package lg

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/nn/nnresponse"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	lg2 "github.com/flipped-aurora/gin-vue-admin/server/utils/lg"
)

type InvoiceService struct {
}

func (invoiceService *InvoiceService) CreateInvoice(invoice lg.Invoice) (err error) {
	err = global.GVA_DB.Create(&invoice).Error
	return err
}

func (invoiceService *InvoiceService) DeleteInvoice(invoice lg.Invoice) (err error) {
	err = global.GVA_DB.Delete(&invoice).Error
	return err
}

func (invoiceService *InvoiceService) DeleteInvoiceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Invoice{}, "id in ?", ids.Ids).Error
	return err
}

func (invoiceService *InvoiceService) UpdateInvoice(invoice lg.Invoice) (err error) {
	err = global.GVA_DB.Save(&invoice).Error
	return err
}

func (invoiceService *InvoiceService) GetInvoice(id uint) (invoice lg.Invoice, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&invoice).Error
	return
}

func (invoiceService *InvoiceService) GetInvoiceInfoList(info lgReq.InvoiceSearch) (list []lg.Invoice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Invoice{})
	var invoices []lg.Invoice
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&invoices).Error
	return invoices, total, err
}

func AutoMaticQueryInvoiceResult() {
	var invoices []lg.Invoice
	err := global.GVA_DB.Where("invoice_download_url is null OR invoice_download_url = ''").Find(&invoices).Error
	if err != nil {
		return
	}
	for _, invoice := range invoices {
		go QueryInvoiceResult(invoice)
	}
}

func QueryInvoiceResult(invoice lg.Invoice) {
	type NNReq struct {
		SerialNos            []string `json:"serialNos,omitempty"`
		OrderNos             []string `json:"orderNos,omitempty"`
		IsOfferInvoiceDetail string   `json:"isOfferInvoiceDetail,omitempty"`
	}
	nnreq := NNReq{
		SerialNos: []string{*invoice.InvoiceSerialNum},
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
		return
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
	invoice.InvoiceDownloadUrl = &resultRequestInvoice.Result[0].PdfUrl
	_ = global.GVA_DB.Save(&invoice).Error
}
