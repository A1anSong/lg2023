package jrclientrequest

type DelayPush struct {
	OrderNo         string  `json:"orderNo"`
	ApplyNo         string  `json:"applyNo"`
	ElogNo          string  `json:"elogNo"`
	AuditStatus     int64   `json:"auditStatus"`
	AuditOpinion    string  `json:"auditOpinion"`
	AuditDate       string  `json:"auditDate"`
	ElogUrl         string  `json:"elogUrl"`
	ElogEncryptUrl  string  `json:"elogEncryptUrl"`
	TenderDeposit   float64 `json:"tenderDeposit"`
	InsureStartDate string  `json:"insureStartDate"`
	InsureEndDate   string  `json:"insureEndDate"`
	InsureDay       int64   `json:"insureDay"`
	ValidateCode    string  `json:"validateCode"`
}
