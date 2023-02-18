package jrrequest

type Order struct {
	OrderNo        *string  `json:"orderNo"`
	OrderPayAmount *float64 `json:"orderPayAmount"`
}
