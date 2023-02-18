package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Delay struct {
	global.GVA_MODEL
	OrderID           *uint    `json:"orderID" form:"orderID"`
	Order             *Order   `json:"-" form:"-"`
	OrderNo           *string  `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	ApplyNo           *string  `json:"applyNo" form:"applyNo" gorm:"type:varchar(64);"`
	ElogNo            *string  `json:"elogNo" form:"elogNo" gorm:"type:varchar(128);"`
	ProjectGuid       *string  `json:"projectGuid" form:"projectGuid" gorm:"type:varchar(64);"`
	ProjectName       *string  `json:"projectName" form:"projectName" gorm:"type:varchar(256);"`
	ProjectNo         *string  `json:"projectNo" form:"projectNo" gorm:"type:varchar(128);"`
	TenderDeposit     *float64 `json:"tenderDeposit" form:"tenderDeposit"`
	DepositStartDate  *string  `json:"depositStartDate" form:"depositStartDate" gorm:"type:varchar(19);"`
	DepositEndDate    *string  `json:"depositEndDate" form:"depositEndDate" gorm:"type:varchar(19);"`
	OpenBeginDate     *string  `json:"openBeginDate" form:"openBeginDate" gorm:"type:varchar(19);"`
	InsureName        *string  `json:"insureName" form:"insureName" gorm:"type:varchar(256);"`
	InsureCreditCode  *string  `json:"insureCreditCode" form:"insureCreditCode" gorm:"type:varchar(18);"`
	ApplicantName     *string  `json:"applicantName" form:"applicantName" gorm:"type:varchar(64);"`
	ApplicantIdCard   *string  `json:"applicantIdCard" form:"applicantIdCard" gorm:"type:varchar(18);"`
	ApplicantTel      *string  `json:"applicantTel" form:"applicantTel" gorm:"type:varchar(11);"`
	Reason            *string  `json:"reason" form:"reason" gorm:"type:text;"`
	AttachInfo        *string  `json:"attachInfo" form:"attachInfo" gorm:"type:text;"`
	AuditStatus       *int64   `json:"auditStatus" form:"auditStatus"`
	AuditOpinion      *string  `json:"auditOpinion" form:"auditOpinion" gorm:"type:varchar(512);"`
	AuditDate         *string  `json:"auditDate" form:"auditDate" gorm:"type:varchar(19);"`
	ElogUrl           *string  `json:"elogUrl" form:"elogUrl" gorm:"type:varchar(256);"`
	ElogFileID        *uint    `json:"elogFileID" form:"elogFileID"`
	ElogFile          *File    `json:"elogFile" form:"elogFile"`
	ElogEncryptUrl    *string  `json:"elogEncryptUrl" form:"elogEncryptUrl" gorm:"type:varchar(256);"`
	ElogEncryptFileID *uint    `json:"elogEncryptFileID" form:"elogEncryptFileID"`
	ElogEncryptFile   *File    `json:"elogEncryptFile" form:"elogEncryptFile"`
	InsureStartDate   *string  `json:"insureStartDate" form:"insureStartDate" gorm:"type:varchar(19);"`
	InsureEndDate     *string  `json:"insureEndDate" form:"insureEndDate" gorm:"type:varchar(19);"`
	InsureDay         *int64   `json:"insureDay" form:"insureDay"`
	ValidateCode      *string  `json:"validateCode" form:"validateCode" gorm:"type:varchar(8);"`
}

func (Delay) TableName() string {
	return "lg_delay"
}
