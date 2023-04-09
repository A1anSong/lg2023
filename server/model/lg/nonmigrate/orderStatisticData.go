package nonmigrate

type OrderStatisticData struct {
	TotalOrderCount          *int     `json:"totalOrderCount" form:"totalOrderCount"`
	TotalGuaranteeAmount     *float64 `json:"totalGuaranteeAmount" form:"totalGuaranteeAmount"`
	TotalElogAmount          *float64 `json:"totalElogAmount" form:"totalElogAmount"`
	TodayOrderCount          *int     `json:"todayOrderCount" form:"todayOrderCount"`
	TodayGuaranteeAmount     *float64 `json:"todayGuaranteeAmount" form:"todayGuaranteeAmount"`
	TodayElogAmount          *float64 `json:"todayElogAmount" form:"todayElogAmount"`
	YesterdayOrderCount      *int     `json:"yesterdayOrderCount" form:"yesterdayOrderCount"`
	YesterdayGuaranteeAmount *float64 `json:"yesterdayGuaranteeAmount" form:"yesterdayGuaranteeAmount"`
	YesterdayElogAmount      *float64 `json:"yesterdayElogAmount" form:"yesterdayElogAmount"`
	WeekOrderCount           *int     `json:"weekOrderCount" form:"weekOrderCount"`
	WeekGuaranteeAmount      *float64 `json:"weekGuaranteeAmount" form:"weekGuaranteeAmount"`
	WeekElogAmount           *float64 `json:"weekElogAmount" form:"weekElogAmount"`
	LastWeekOrderCount       *int     `json:"lastWeekOrderCount" form:"lastWeekOrderCount"`
	LastWeekGuaranteeAmount  *float64 `json:"lastWeekGuaranteeAmount" form:"lastWeekGuaranteeAmount"`
	LastWeekElogAmount       *float64 `json:"lastWeekElogAmount" form:"lastWeekElogAmount"`
	MonthOrderCount          *int     `json:"monthOrderCount" form:"monthOrderCount"`
	MonthGuaranteeAmount     *float64 `json:"monthGuaranteeAmount" form:"monthGuaranteeAmount"`
	MonthElogAmount          *float64 `json:"monthElogAmount" form:"monthElogAmount"`
	LastMonthOrderCount      *int     `json:"lastMonthOrderCount" form:"lastMonthOrderCount"`
	LastMonthGuaranteeAmount *float64 `json:"lastMonthGuaranteeAmount" form:"lastMonthGuaranteeAmount"`
	LastMonthElogAmount      *float64 `json:"lastMonthElogAmount" form:"lastMonthElogAmount"`
}
