package jrresponse

type JRAPIApply struct {
	OrderNo      *string `json:"orderNo"`
	ApplyNo      *string `json:"applyNo"`
	AuditStatus  *int64  `json:"auditStatus"`
	AuditOpinion *string `json:"auditOpinion"`
	AuditDate    *string `json:"auditDate"`
}
