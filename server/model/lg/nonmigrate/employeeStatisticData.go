package nonmigrate

type EmployeeStatisticData struct {
	EmployeeData []EmployeeTrendDataItem `json:"employeeData" form:"employeeData"`
}

type EmployeeTrendDataItem struct {
	Name  *string `json:"name" form:"name"`
	Count *int    `json:"count" form:"count"`
}
