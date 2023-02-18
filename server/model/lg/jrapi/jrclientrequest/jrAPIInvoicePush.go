package jrclientrequest

type InvoicePush struct {
	ApplyNo            string    `json:"applyNo"`
	AuditStatus        int64     `json:"auditStatus"`
	AuditOpinion       string    `json:"auditOpinion"`
	AuditDate          string    `json:"auditDate"`
	InvoiceTotalAmount float64   `json:"invoiceTotalAmount,omitempty"`
	InvoiceList        []Invoice `json:"invoiceList,omitempty"`
}
