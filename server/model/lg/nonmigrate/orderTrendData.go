package nonmigrate

type OrderTrendData struct {
	Seven  []TrendDataItem `json:"seven" form:"seven"`
	Thirty []TrendDataItem `json:"thirty" form:"thirty"`
}

type TrendDataItem struct {
	Date            *string  `json:"date" form:"date"`
	OrderCount      *int     `json:"orderCount" form:"orderCount"`
	GuaranteeAmount *float64 `json:"guaranteeAmount" form:"guaranteeAmount"`
}
