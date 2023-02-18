package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type File struct {
	global.GVA_MODEL
	FileName  *string `json:"fileName" form:"fileName" gorm:"type:varchar(128);"`
	FileSteam []byte  `json:"-" form:"-" gorm:"type:mediumblob"`
}

func (File) TableName() string {
	return "lg_file"
}
