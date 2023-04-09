package nonmigrate

type OrderTrendData struct {
	TrendData7Days  []TrendDataItem `json:"trendData7Days" form:"trendData7Days"`
	TrendData30Days []TrendDataItem `json:"trendData30Days" form:"trendData30Days"`
}

type TrendDataItem struct {
	Date            *string  `json:"date" form:"date"`
	OrderCount      *int     `json:"orderCount" form:"orderCount"`
	GuaranteeAmount *float64 `json:"guaranteeAmount" form:"guaranteeAmount"`
	ElogAmount      *float64 `json:"elogAmount" form:"elogAmount"`
}
