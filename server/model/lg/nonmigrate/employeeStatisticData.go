package nonmigrate

type EmployeeStatisticData struct {
	Day   []EmployeeStatisticDataItem `json:"day" form:"day"`
	Week  []EmployeeStatisticDataItem `json:"week" form:"week"`
	Month []EmployeeStatisticDataItem `json:"month" form:"month"`
	Total []EmployeeStatisticDataItem `json:"total" form:"total"`
}

type EmployeeStatisticDataItem struct {
	Name  *string `json:"name" form:"name"`
	Count *int    `json:"count" form:"count"`
}
