package request

type AssignOrder struct {
	OrderNo    string `json:"orderNo" form:"orderNo"`
	EmployeeId uint   `json:"employeeId" form:"employeeId"`
}
