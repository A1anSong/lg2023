package jrrequest

type JRAPIApplyInvoice struct {
	ApplyNo            *string  `json:"applyNo"`
	InvoiceTotalAmount *float64 `json:"invoiceTotalAmount"`
	InvoiceType        *string  `json:"invoiceType"`
	InvoiceTileType    *string  `json:"invoiceTileType"`
	InvoiceTile        *string  `json:"invoiceTile"`
	TaxNo              *string  `json:"taxNo"`
	BankName           *string  `json:"bankName"`
	BankNo             *string  `json:"bankNo"`
	CompanyAddress     *string  `json:"companyAddress"`
	CompanyTel         *string  `json:"companyTel"`
	Remarks            *string  `json:"remarks"`
	OrderList          []Order  `json:"orderList"`
}
