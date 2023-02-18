package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Claim struct {
	global.GVA_MODEL
	OrderID           *uint    `json:"orderID" form:"orderID"`
	Order             *Order   `json:"-" form:"-"`
	OrderNo           *string  `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	ApplyNo           *string  `json:"applyNo" form:"applyNo" gorm:"type:varchar(64);"`
	ElogNo            *string  `json:"elogNo" form:"elogNo" gorm:"type:varchar(128);"`
	InsuredName       *string  `json:"insuredName" form:"insuredName" gorm:"type:varchar(256);"`
	InsuredCreditCode *string  `json:"insuredCreditCode" form:"insuredCreditCode" gorm:"type:varchar(18);"`
	InsuredBankNo     *string  `json:"insuredBankNo" form:"insuredBankNo" gorm:"type:varchar(19);"`
	InsuredBankName   *string  `json:"insuredBankName" form:"insuredBankName" gorm:"type:varchar(256);"`
	ApplicantName     *string  `json:"applicantName" form:"applicantName" gorm:"type:varchar(64);"`
	ApplicantIdCard   *string  `json:"applicantIdCard" form:"applicantIdCard" gorm:"type:varchar(18);"`
	ApplicantTel      *string  `json:"applicantTel" form:"applicantTel" gorm:"type:varchar(11);"`
	ClaimAmount       *float64 `json:"claimAmount" form:"claimAmount"`
	Reason            *string  `json:"reason" form:"reason" gorm:"type:text;"`
	AttachInfo        *string  `json:"attachInfo" form:"attachInfo" gorm:"type:text;"`
	AuditStatus       *int64   `json:"auditStatus" form:"auditStatus"`
	AuditOpinion      *string  `json:"auditOpinion" form:"auditOpinion" gorm:"type:varchar(512);"`
	AuditDate         *string  `json:"auditDate" form:"auditDate" gorm:"type:varchar(19);"`
	RealClaimAmount   *float64 `json:"realClaimAmount" form:"realClaimAmount"`
	RealClaimDate     *string  `json:"realClaimDate" form:"realClaimDate" gorm:"type:varchar(19);"`
}

func (Claim) TableName() string {
	return "lg_claim"
}
