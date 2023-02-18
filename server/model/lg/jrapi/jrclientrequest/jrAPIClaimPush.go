package jrclientrequest

type ClaimPush struct {
	OrderNo         string  `json:"orderNo"`
	ApplyNo         string  `json:"applyNo"`
	ElogNo          string  `json:"elogNo"`
	AuditStatus     int64   `json:"auditStatus"`
	AuditOpinion    string  `json:"auditOpinion"`
	AuditDate       string  `json:"auditDate"`
	RealClaimAmount float64 `json:"realClaimAmount,omitempty"`
	RealClaimDate   string  `json:"realClaimDate,omitempty"`
}
