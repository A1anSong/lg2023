package nonmigrate

type InsuranceBalance struct {
	InsuranceBalance *float64 `json:"insuranceBalance" form:"insuranceBalance"`
	InsuranceTotal   *float64 `json:"insuranceTotal" form:"insuranceTotal"`
}
