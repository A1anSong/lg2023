package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Apply struct {
	global.GVA_MODEL
	OrderID             *uint    `json:"orderID" form:"orderID"`
	Order               *Order   `json:"-" form:"-"`
	OrderNo             *string  `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	ApplyNo             *string  `json:"applyNo" form:"applyNo" gorm:"type:varchar(64);"`
	ProductNo           *string  `json:"productNo" form:"productNo" gorm:"type:varchar(64);"`
	ProductType         *int64   `json:"productType" form:"productType"`
	ProductRate         *float64 `json:"productRate" form:"productRate"`
	ElogAmount          *float64 `json:"elogAmount" form:"elogAmount"`
	ProjectGuid         *string  `json:"projectGuid" form:"projectGuid" gorm:"type:varchar(64);"`
	ProjectName         *string  `json:"projectName" form:"projectName" gorm:"type:varchar(256);"`
	ProjectNo           *string  `json:"projectNo" form:"projectNo" gorm:"type:varchar(128);"`
	TenderDeposit       *float64 `json:"tenderDeposit" form:"type:tenderDeposit"`
	DepositStartDate    *string  `json:"depositStartDate" form:"depositStartDate" gorm:"type:varchar(19);"`
	DepositEndDate      *string  `json:"depositEndDate" form:"depositEndDate" gorm:"type:varchar(19);"`
	OpenBeginDate       *string  `json:"openBeginDate" form:"openBeginDate" gorm:"type:varchar(19);"`
	ElogTemplateNo      *string  `json:"elogTemplateNo" form:"elogTemplateNo" gorm:"type:varchar(14);"`
	ElogTemplateName    *string  `json:"elogTemplateName" form:"elogTemplateName" gorm:"type:varchar(256);"`
	InsuredName         *string  `json:"insuredName" form:"insuredName" gorm:"type:varchar(256);"`
	InsuredCreditCode   *string  `json:"insuredCreditCode" form:"insuredCreditCode" gorm:"type:varchar(64);"`
	InsuredAddress      *string  `json:"insuredAddress" form:"insuredAddress" gorm:"type:varchar(256);"`
	InsureName          *string  `json:"insureName" form:"insureName" gorm:"type:varchar(256);"`
	InsureCreditCode    *string  `json:"insureCreditCode" form:"insureCreditCode" gorm:"type:varchar(64);"`
	InsureLegalName     *string  `json:"insureLegalName" form:"insureLegalName" gorm:"type:varchar(64);"`
	InsureLegalIdCard   *string  `json:"insureLegalIdCard" form:"insureLegalIdCard" gorm:"type:varchar(64);"`
	InsureAddress       *string  `json:"insureAddress" form:"insureAddress" gorm:"type:varchar(256);"`
	ApplicantName       *string  `json:"applicantName" form:"applicantName" gorm:"type:varchar(64);"`
	ApplicantIdCard     *string  `json:"applicantIdCard" form:"applicantIdCard" gorm:"type:varchar(64);"`
	ApplicantTel        *string  `json:"applicantTel" form:"applicantTel" gorm:"type:varchar(64);"`
	ApplicantAuthCode   *string  `json:"applicantAuthCode" form:"applicantAuthCode" gorm:"type:varchar(12);"`
	AttachInfo          *string  `json:"attachInfo" form:"attachInfo" gorm:"type:text;"`
	AuditStatus         *int64   `json:"auditStatus" form:"auditStatus"`
	AuditOpinion        *string  `json:"auditOpinion" form:"auditOpinion" gorm:"type:varchar(512);"`
	AuditDate           *string  `json:"auditDate" form:"auditDate" gorm:"type:varchar(19);"`
	RealElogAmount      *float64 `json:"realElogAmount" form:"realElogAmount"`
	RealElogRate        *float64 `json:"realElogRate" form:"realElogRate"`
	InsuranceName       *string  `json:"insuranceName" form:"insuranceName" gorm:"type:varchar(256);"`
	InsuranceCreditCode *string  `json:"insuranceCreditCode" form:"insuranceCreditCode" gorm:"type:varchar(18);"`
	PayUrl              *string  `json:"payUrl" form:"payUrl" gorm:"type:varchar(256);"`
	ReceiveName         *string  `json:"receiveName" form:"receiveName" gorm:"type:varchar(256);"`
}

func (Apply) TableName() string {
	return "lg_apply"
}
