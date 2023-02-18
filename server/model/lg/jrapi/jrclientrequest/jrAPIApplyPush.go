package jrclientrequest

type ApplyPush struct {
	OrderNo             string  `json:"orderNo"`
	ApplyNo             string  `json:"applyNo"`
	AuditStatus         int64   `json:"auditStatus"`
	AuditOpinion        string  `json:"auditOpinion"`
	AuditDate           string  `json:"auditDate"`
	RealElogAmount      float64 `json:"realElogAmount"`
	RealElogRate        float64 `json:"realElogRate"`
	TenderDeposit       float64 `json:"tenderDeposit"`
	InsuranceName       string  `json:"insuranceName"`
	InsuranceCreditCode string  `json:"insuranceCreditCode"`
	PayUrl              string  `json:"payUrl,omitempty"`
	ReceiveName         string  `json:"receiveName,omitempty"`
}
