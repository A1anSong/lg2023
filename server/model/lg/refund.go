package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Refund struct {
	global.GVA_MODEL
	OrderID          *uint    `json:"orderID" form:"orderID"`
	Order            *Order   `json:"-" form:"-"`
	OrderNo          *string  `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	ApplyNo          *string  `json:"applyNo" form:"applyNo" gorm:"type:varchar(64);"`
	ElogNo           *string  `json:"elogNo" form:"elogNo" gorm:"type:varchar(128);"`
	InsureName       *string  `json:"insureName" form:"insureName" gorm:"type:varchar(256);"`
	InsureCreditCode *string  `json:"insureCreditCode" form:"insureCreditCode" gorm:"type:varchar(18);"`
	ApplicantName    *string  `json:"applicantName" form:"applicantName" gorm:"type:varchar(64);"`
	ApplicantIdCard  *string  `json:"applicantIdCard" form:"applicantIdCard" gorm:"type:varchar(18);"`
	ApplicantTel     *string  `json:"applicantTel" form:"applicantTel" gorm:"type:varchar(64);"`
	Reason           *string  `json:"reason" form:"reason" gorm:"type:text;"`
	AttachInfo       *string  `json:"attachInfo" form:"attachInfo" gorm:"type:text;"`
	AuditStatus      *int64   `json:"auditStatus" form:"auditStatus"`
	AuditOpinion     *string  `json:"auditOpinion" form:"auditOpinion" gorm:"type:varchar(512);"`
	AuditDate        *string  `json:"auditDate" form:"auditDate" gorm:"type:varchar(19);"`
	PayAmount        *float64 `json:"payAmount" form:"payAmount"`
}

func (Refund) TableName() string {
	return "lg_refund"
}
