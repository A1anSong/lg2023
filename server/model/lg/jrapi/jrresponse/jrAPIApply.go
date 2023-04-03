package jrresponse

type JRAPIApply struct {
	OrderNo             *string  `json:"orderNo"`
	ApplyNo             *string  `json:"applyNo"`
	AuditStatus         *int64   `json:"auditStatus"`
	AuditOpinion        *string  `json:"auditOpinion"`
	AuditDate           *string  `json:"auditDate"`
	RealElogAmount      *float64 `json:"realElogAmount,omitempty"`
	RealElogRate        *float64 `json:"realElogRate,omitempty"`
	TenderDeposit       *float64 `json:"tenderDeposit,omitempty"`
	InsuranceName       *string  `json:"insuranceName,omitempty"`
	InsuranceCreditCode *string  `json:"insuranceCreditCode,omitempty"`
}
