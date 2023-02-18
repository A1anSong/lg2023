package jrclientrequest

type Order struct {
	OrderNo            string  `json:"orderNo"`
	OrderInvoiceAmount float64 `json:"orderInvoiceAmount"`
}
