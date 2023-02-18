package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/nonmigrate"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"os"
)

type TemplateService struct {
}

func (templateService *TemplateService) CreateTemplate(templateAndFile nonmigrate.TemplateAndFile) (err error) {
	//err = global.GVA_DB.Create(&template).Error
	basePath := "./tmp/"
	file, err := os.Open(basePath + *templateAndFile.FileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	templateAndFile.FileSteam = fileContent

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		file := templateAndFile.File
		if err = tx.Create(&file).Error; err != nil {
			return err
		}
		template := templateAndFile.Template
		template.TemplateFileID = &file.ID
		if err = tx.Create(&template).Error; err != nil {
			return err
		}
		_ = os.Remove(basePath + *templateAndFile.FileName)
		return nil
	})

	return err
}

func (templateService *TemplateService) DeleteTemplate(template lg.Template) (err error) {
	err = global.GVA_DB.Select(clause.Associations).Delete(&template).Error
	return err
}

func (templateService *TemplateService) DeleteTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Template{}, "id in ?", ids.Ids).Error
	return err
}

func (templateService *TemplateService) UpdateTemplate(template lg.Template) (err error) {
	err = global.GVA_DB.Save(&template).Error
	return err
}

func (templateService *TemplateService) GetTemplate(id uint) (template lg.Template, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&template).Error
	return
}

func (templateService *TemplateService) GetTemplateInfoList(info lgReq.TemplateSearch) (list []lg.Template, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Template{})
	var templates []lg.Template
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Preload(clause.Associations).Order("created_at desc").Offset(offset).Find(&templates).Error
	return templates, total, err
}
