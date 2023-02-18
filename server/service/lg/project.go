package lg

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	"gorm.io/gorm"
)

type ProjectService struct {
}

func (projectService *ProjectService) CreateProject(project lg.Project) (err error) {
	fmt.Println(project.ID)
	err = global.GVA_DB.Create(&project).Error
	fmt.Println(project.ID)
	return err
}

func (projectService *ProjectService) DeleteProject(project lg.Project) (err error) {
	err = global.GVA_DB.Delete(&project).Error
	return err
}

func (projectService *ProjectService) DeleteProjectByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Project{}, "id in ?", ids.Ids).Error
	return err
}

func (projectService *ProjectService) UpdateProject(project lg.Project) (err error) {
	err = global.GVA_DB.Save(&project).Error
	return err
}

func (projectService *ProjectService) GetProject(id uint) (project lg.Project, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&project).Error
	return
}

func (projectService *ProjectService) GetProjectInfoList(info lgReq.ProjectSearch) (list []lg.Project, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Project{})
	var projects []lg.Project
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Order("created_at desc").Offset(offset).Find(&projects).Error
	return projects, total, err
}

func (projectService *ProjectService) BindProject(project lg.Project) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var orders []lg.Order
		err = tx.Model(&lg.Order{}).Joins("Apply").Where("Apply.project_no = ?", project.ProjectNo).Find(&orders).Error
		if err != nil {
			return err
		}
		if len(orders) > 0 {
			for i := range orders {
				orders[i].ProjectID = &project.ID
			}
			err = tx.Save(&orders).Error
			if err != nil {
				return err
			}
		}
		err = tx.Save(&project).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (projectService *ProjectService) UnbindProject(project lg.Project) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var orders []lg.Order
		err = tx.Model(&lg.Order{}).Joins("Apply").Where("Apply.project_no = ?", project.ProjectNo).Find(&orders).Error
		if err != nil {
			return err
		}
		if len(orders) > 0 {
			for i := range orders {
				orders[i].ProjectID = nil
			}
			err = tx.Save(&orders).Error
			if err != nil {
				return err
			}
		}
		err = tx.Save(&project).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
