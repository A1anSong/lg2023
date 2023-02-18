package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Pay struct {
	global.GVA_MODEL
	OrderID    *uint    `json:"orderID" form:"orderID"`
	Order      *Order   `json:"-" form:"-"`
	OrderNo    *string  `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	PayNo      *string  `json:"payNo" form:"payNo" gorm:"type:varchar(64);"`
	PayAmount  *float64 `json:"payAmount" form:"payAmount"`
	PayTime    *string  `json:"payTime" form:"payTime" gorm:"type:varchar(19);"`
	PayTransNo *string  `json:"payTransNo" form:"payTransNo" gorm:"type:varchar(256);"`
}

func (Pay) TableName() string {
	return "lg_pay"
}
