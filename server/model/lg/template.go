package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Template struct {
	global.GVA_MODEL
	TemplateName   *string `json:"templateName" form:"templateName" gorm:"type:varchar(128);"`
	TemplateFileID *uint   `json:"templateFileID" form:"templateFileID"`
	TemplateFile   *File   `json:"templateFile" form:"templateFile"`
}

func (Template) TableName() string {
	return "lg_template"
}
