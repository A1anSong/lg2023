package jrrequest

type JRAPIPayPush struct {
	OrderNo    *string  `json:"orderNo"`
	PayNo      *string  `json:"payNo"`
	PayAmount  *float64 `json:"payAmount"`
	PayTime    *string  `json:"payTime"`
	PayTransNo *string  `json:"payTransNo"`
}
