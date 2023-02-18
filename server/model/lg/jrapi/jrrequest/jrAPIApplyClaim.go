package jrrequest

type JRAPIApplyClaim struct {
	OrderID           *uint             `json:"orderID"`
	OrderNo           *string           `json:"orderNo"`
	ApplyNo           *string           `json:"applyNo"`
	ElogNo            *string           `json:"elogNo"`
	InsuredName       *string           `json:"insuredName"`
	InsuredCreditCode *string           `json:"insuredCreditCode"`
	InsuredBankNo     *string           `json:"insuredBankNo"`
	InsuredBankName   *string           `json:"insuredBankName"`
	ApplicantName     *string           `json:"applicantName"`
	ApplicantIdCard   *string           `json:"applicantIdCard"`
	ApplicantTel      *string           `json:"applicantTel"`
	ClaimAmount       *float64          `json:"claimAmount"`
	Reason            *string           `json:"reason"`
	AttachInfo        []JRAPIAttachInfo `json:"attachInfo"`
}
