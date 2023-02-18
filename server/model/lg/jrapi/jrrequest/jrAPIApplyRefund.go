package jrrequest

type JRAPIApplyRefund struct {
	OrderNo          *string           `json:"orderNo"`
	ApplyNo          *string           `json:"applyNo"`
	ElogNo           *string           `json:"elogNo"`
	InsureName       *string           `json:"insureName"`
	InsureCreditCode *string           `json:"insureCreditCode"`
	ApplicantName    *string           `json:"applicantName"`
	ApplicantIdCard  *string           `json:"applicantIdCard"`
	ApplicantTel     *string           `json:"applicantTel"`
	Reason           *string           `json:"reason"`
	AttachInfo       []JRAPIAttachInfo `json:"attachInfo"`
}
