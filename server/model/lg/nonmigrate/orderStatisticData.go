package nonmigrate

type OrderStatisticData struct {
	Day       OrderStatisticDataItem `json:"day" form:"day"`
	Lastday   OrderStatisticDataItem `json:"lastday" form:"lastday"`
	Week      OrderStatisticDataItem `json:"week" form:"week"`
	Lastweek  OrderStatisticDataItem `json:"lastweek" form:"lastweek"`
	Month     OrderStatisticDataItem `json:"month" form:"month"`
	Lastmonth OrderStatisticDataItem `json:"lastmonth" form:"lastmonth"`
	Total     OrderStatisticDataItem `json:"total" form:"total"`
}

type OrderStatisticDataItem struct {
	OrderCount      *int     `json:"orderCount" form:"orderCount"`
	GuaranteeAmount *float64 `json:"guaranteeAmount" form:"guaranteeAmount"`
	ElogAmount      *float64 `json:"elogAmount" form:"elogAmount"`
}
