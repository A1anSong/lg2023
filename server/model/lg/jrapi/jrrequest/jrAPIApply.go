package jrrequest

type JRAPIApply struct {
	OrderNo           *string           `json:"orderNo"`
	ApplyNo           *string           `json:"applyNo"`
	ProductNo         *string           `json:"productNo"`
	ProductType       *string           `json:"productType"`
	ProductRate       *float64          `json:"productRate"`
	ElogAmount        *float64          `json:"elogAmount"`
	ProjectGuid       *string           `json:"projectGuid"`
	ProjectName       *string           `json:"projectName"`
	ProjectNo         *string           `json:"projectNo"`
	TenderDeposit     *float64          `json:"tenderDeposit"`
	DepositStartDate  *string           `json:"depositStartDate"`
	DepositEndDate    *string           `json:"depositEndDate"`
	OpenBeginDate     *string           `json:"openBeginDate"`
	ElogTemplateNo    *string           `json:"elogTemplateNo"`
	ElogTemplateName  *string           `json:"elogTemplateName"`
	InsuredName       *string           `json:"insuredName"`
	InsuredCreditCode *string           `json:"insuredCreditCode"`
	InsuredAddress    *string           `json:"insuredAddress"`
	InsureName        *string           `json:"insureName"`
	InsureCreditCode  *string           `json:"insureCreditCode"`
	InsureLegalName   *string           `json:"insureLegalName"`
	InsureLegalIdCard *string           `json:"insureLegalIdCard"`
	InsureAddress     *string           `json:"insureAddress"`
	ApplicantName     *string           `json:"applicantName"`
	ApplicantIdCard   *string           `json:"applicantIdCard"`
	ApplicantTel      *string           `json:"applicantTel"`
	ApplicantAuthCode *string           `json:"applicantAuthCode"`
	AttachInfo        []JRAPIAttachInfo `json:"attachInfo"`
}
