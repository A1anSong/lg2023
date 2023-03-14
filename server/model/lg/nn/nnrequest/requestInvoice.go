package nnrequest

import "github.com/flipped-aurora/gin-vue-admin/server/model/lg"

type NNRequestInvoice struct {
	Order        lg.Order        `json:"order"`
	InvoiceApply lg.InvoiceApply `json:"invoiceApply"`
}
