package jrclientrequest

type Invoice struct {
	InvoiceNo          string  `json:"invoiceNo"`
	InvoiceAmount      float64 `json:"invoiceAmount"`
	InvoiceType        string  `json:"invoiceType"`
	InvoiceForm        string  `json:"invoiceForm"`
	InvoicePoint       float64 `json:"invoicePoint"`
	InvoiceContent     string  `json:"invoiceContent"`
	InvoiceRemark      string  `json:"invoiceRemark,omitempty"`
	InvoiceTime        string  `json:"invoiceTime"`
	InvoiceDownloadUrl string  `json:"invoiceDownloadUrl"`
	OrderList          []Order `json:"orderList"`
}
