package jrclientrequest

type RefundPush struct {
	OrderNo      string  `json:"orderNo"`
	ApplyNo      string  `json:"applyNo"`
	ElogNo       string  `json:"elogNo"`
	AuditStatus  int64   `json:"auditStatus"`
	AuditOpinion string  `json:"auditOpinion"`
	AuditDate    string  `json:"auditDate"`
	PayAmount    float64 `json:"payAmount,omitempty"`
}
