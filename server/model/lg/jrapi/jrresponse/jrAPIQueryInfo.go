package jrresponse

type JRAPIQueryInfo struct {
	OrderNo             *string  `json:"orderNo"`
	ElogNo              *string  `json:"elogNo"`
	ProductNo           *string  `json:"productNo"`
	ProductType         *int64   `json:"productType"`
	ProductRate         *float64 `json:"productRate"`
	ElogAmount          *float64 `json:"elogAmount"`
	InsuranceName       *string  `json:"insuranceName"`
	InsuranceCreditCode *string  `json:"insuranceCreditCode"`
	ElogOutDate         *string  `json:"elogOutDate"`
	ElogUrl             *string  `json:"elogUrl"`
	ElogEncryptUrl      *string  `json:"elogEncryptUrl"`
	TenderDeposit       *float64 `json:"tenderDeposit"`
	InsureStartDate     *string  `json:"insureStartDate"`
	InsureEndDate       *string  `json:"insureEndDate"`
	InsureDay           *int64   `json:"insureDay"`
	ValidateCode        *string  `json:"validateCode"`
}
