package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Logout struct {
	global.GVA_MODEL
	ProjectGuid         *string `json:"projectGuid" form:"projectGuid" gorm:"type:varchar(64);"`
	ProjectName         *string `json:"projectName" form:"projectName" gorm:"type:varchar(256);"`
	ProjectNo           *string `json:"projectNo" form:"projectNo" gorm:"type:varchar(128);"`
	Reason              *string `json:"reason" form:"reason" gorm:"type:text;"`
	LogoutType          *int64  `json:"logoutType" form:"logoutType"`
	WinBidderName       *string `json:"winBidderName" form:"winBidderName" gorm:"type:varchar(256);"`
	WinBidderCreditCode *string `json:"winBidderCreditCode" form:"winBidderCreditCode" gorm:"type:varchar(18);"`
}

func (Logout) TableName() string {
	return "lg_logout"
}
