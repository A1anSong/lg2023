package lg

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/nn/nnresponse"
	lg2 "github.com/flipped-aurora/gin-vue-admin/server/utils/lg"
)

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
