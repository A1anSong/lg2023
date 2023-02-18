package jrrequest

type JRAPIApplyDelay struct {
	OrderNo          *string           `json:"orderNo"`
	ApplyNo          *string           `json:"applyNo"`
	ElogNo           *string           `json:"elogNo"`
	ProjectGuid      *string           `json:"projectGuid"`
	ProjectName      *string           `json:"projectName"`
	ProjectNo        *string           `json:"projectNo"`
	TenderDeposit    *float64          `json:"tenderDeposit"`
	DepositStartDate *string           `json:"depositStartDate"`
	DepositEndDate   *string           `json:"depositEndDate"`
	OpenBeginDate    *string           `json:"openBeginDate"`
	InsureName       *string           `json:"insureName"`
	InsureCreditCode *string           `json:"insureCreditCode"`
	ApplicantName    *string           `json:"applicantName"`
	ApplicantIdCard  *string           `json:"applicantIdCard"`
	ApplicantTel     *string           `json:"applicantTel"`
	Reason           *string           `json:"reason"`
	AttachInfo       []JRAPIAttachInfo `json:"attachInfo"`
}
