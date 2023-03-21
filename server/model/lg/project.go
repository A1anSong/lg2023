package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Project struct {
	global.GVA_MODEL
	ProjectName        *string   `json:"projectName" form:"projectName" gorm:"type:varchar(256);"`
	ProjectNo          *string   `json:"projectNo" form:"projectNo" gorm:"type:varchar(64);"`
	ProjectAmount      *float64  `json:"projectAmount" form:"projectAmount"`
	TendereeName       *string   `json:"tendereeName" form:"tendereeName" gorm:"type:varchar(256);"`
	TendereeAddress    *string   `json:"tendereeAddress" form:"tendereeAddress" gorm:"type:varchar(256);"`
	TendereeTel        *string   `json:"tendereeTel" form:"tendereeTel" gorm:"type:varchar(64);"`
	TendereeFile       *string   `json:"tendereeFile" form:"tendereeFile" gorm:"type:varchar(512)"`
	AgentTel           *string   `json:"agentTel" form:"agentTel" gorm:"type:varchar(64)"`
	TenderDeposit      *float64  `json:"tenderDeposit" form:"tenderDeposit"`
	ProjectOpenTime    *string   `json:"projectOpenTime" form:"projectOpenTime" gorm:"type:varchar(19);"`
	ProjectPublishTime *string   `json:"projectPublishTime" form:"projectPublishTime" gorm:"type:varchar(19);"`
	ProjectCity        *string   `json:"projectCity" form:"projectCity" gorm:"type:varchar(64);"`
	ProjectCounty      *string   `json:"projectCounty" form:"projectCounty" gorm:"type:varchar(64);"`
	ProjectDay         *int64    `json:"projectDay" form:"projectDay"`
	TenderEndDate      *string   `json:"tenderEndDate" form:"tenderEndDate" gorm:"type:varchar(19);"`
	ProjectType        *string   `json:"projectType" form:"projectType" gorm:"type:varchar(64);"`
	ProjectCategory    *string   `json:"projectCategory" form:"projectCategory" gorm:"type:varchar(64);"`
	TemplateID         *uint     `json:"templateID" form:"templateID"`
	Template           *Template `json:"template" form:"template"`
	IsEnable           *bool     `json:"isEnable" form:"isEnable"`
}

func (Project) TableName() string {
	return "lg_project"
}
