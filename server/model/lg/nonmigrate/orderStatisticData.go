package nonmigrate

type OrderStatisticData struct {
	TotalGuaranteeAmount *float64 `json:"totalGuaranteeAmount" form:"totalGuaranteeAmount"`
	TotalElogAmount      *float64 `json:"totalElogAmount" form:"totalElogAmount"`
}
