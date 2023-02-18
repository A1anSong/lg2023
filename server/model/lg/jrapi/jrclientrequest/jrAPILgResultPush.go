package jrclientrequest

type LgResultPush struct {
	OrderNo             string  `json:"orderNo"`
	ElogNo              string  `json:"elogNo"`
	InsuranceName       string  `json:"insuranceName"`
	InsuranceCreditCode string  `json:"insuranceCreditCode"`
	ElogOutDate         string  `json:"elogOutDate"`
	ElogUrl             string  `json:"elogUrl"`
	ElogEncryptUrl      string  `json:"elogEncryptUrl"`
	TenderDeposit       float64 `json:"tenderDeposit"`
	InsureStartDate     string  `json:"insureStartDate"`
	InsureEndDate       string  `json:"insureEndDate"`
	InsureDay           int64   `json:"insureDay"`
	ValidateCode        string  `json:"validateCode"`
}
