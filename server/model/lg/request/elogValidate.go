package request

type ElogValidate struct {
	ElogNo       *string `json:"elogNo" form:"elogNo"`
	ValidateCode *string `json:"validateCode" form:"validateCode"`
}
