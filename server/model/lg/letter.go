package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Letter struct {
	global.GVA_MODEL
	OrderID             *uint    `json:"orderID" form:"orderID"`
	Order               *Order   `json:"-" form:"-"`
	OrderNo             *string  `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	ElogNo              *string  `json:"elogNo" form:"elogNo" gorm:"type:varchar(128);"`
	InsuranceName       *string  `json:"insuranceName" form:"insuranceName" gorm:"type:varchar(256);"`
	InsuranceCreditCode *string  `json:"insuranceCreditCode" form:"insuranceCreditCode" gorm:"type:varchar(18);"`
	ElogOutDate         *string  `json:"elogOutDate" form:"elogOutDate" gorm:"type:varchar(19);"`
	ElogUrl             *string  `json:"elogUrl" form:"elogUrl" gorm:"type:varchar(256);"`
	ElogFileID          *uint    `json:"elogFileID" form:"elogFileID"`
	ElogFile            *File    `json:"elogFile" form:"elogFile"`
	ElogEncryptUrl      *string  `json:"elogEncryptUrl" form:"elogEncryptUrl" gorm:"type:varchar(256);"`
	ElogEncryptFileID   *uint    `json:"elogEncryptFileID" form:"elogEncryptFileID"`
	ElogEncryptFile     *File    `json:"elogEncryptFile" form:"elogEncryptFile"`
	TenderDeposit       *float64 `json:"tenderDeposit" form:"tenderDeposit"`
	InsureStartDate     *string  `json:"insureStartDate" form:"insureStartDate" gorm:"type:varchar(19);"`
	InsureEndDate       *string  `json:"insureEndDate" form:"insureEndDate" gorm:"type:varchar(19);"`
	InsureDay           *int64   `json:"insureDay" form:"insureDay"`
	ValidateCode        *string  `json:"validateCode" form:"validateCode" gorm:"type:varchar(8);"`
}

func (Letter) TableName() string {
	return "lg_letter"
}
