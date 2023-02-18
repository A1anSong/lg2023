package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Revoke struct {
	global.GVA_MODEL
	OrderID      *uint   `json:"orderID" form:"orderID"`
	Order        *Order  `json:"-" form:"-"`
	OrderNo      *string `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	RevokeOrigin *string `json:"revokeOrigin" form:"revokeOrigin" gorm:"type:varchar(2);"`
	RevokeReason *string `json:"revokeReason" form:"revokeOrigin" gorm:"type:text;"`
}

func (Revoke) TableName() string {
	return "lg_revoke"
}
