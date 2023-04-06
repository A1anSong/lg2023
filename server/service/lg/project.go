package lg

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type ProjectService struct {
}

func (projectService *ProjectService) CreateProject(project lg.Project) (err error) {
	isEnable := false
	isAutoMatic := false
	project.IsEnable = &isEnable
	project.IsAutoMatic = &isAutoMatic
	err = global.GVA_DB.Create(&project).Error
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
	if info.ProjectNo != nil {
		db = db.Where("project_no = ?", info.ProjectNo)
	}
	if info.ProjectName != nil {
		db = db.Where("project_name = ?", info.ProjectName)
	}
	if info.OpenTime != nil {
		db = db.Where("project_open_time BETWEEN ? AND ?", info.OpenTime[0], info.OpenTime[1])
	}
	if info.ProjectCity != nil {
		db = db.Where("project_city = ?", info.ProjectCity)
	}
	if info.ProjectCounty != nil {
		db = db.Where("project_county = ?", info.ProjectCounty)
	}
	if info.TemplateID != nil {
		db = db.Where("template_id = ?", info.TemplateID)
	}
	if info.IsEnable != nil {
		db = db.Where("is_enable = ?", info.IsEnable)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Debug().Limit(limit).Order("created_at desc").Offset(offset).Find(&projects).Error
	return projects, total, err
}

func (projectService *ProjectService) BindProject(project lg.Project) (err error) {
	if project.TemplateID == nil {
		return errors.New("请先选择模板")
	}
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
	err = global.GVA_DB.Save(&project).Error
	return err
}

func (projectService *ProjectService) AutoMaticProject(project lg.Project) (err error) {
	err = global.GVA_DB.Model(&project).Update("is_auto_matic", true).Error
	return err
}

func (projectService *ProjectService) UnAutoMaticProject(project lg.Project) (err error) {
	err = global.GVA_DB.Model(&project).Update("is_auto_matic", false).Error
	return err
}

func (projectService *ProjectService) ImportExcel(file *multipart.FileHeader) (err error) {
	basePath := "./tmp/"
	fileSuffix := path.Ext(file.Filename)
	fileNameOnly := strings.TrimSuffix(file.Filename, fileSuffix)
	timestampString := strconv.FormatInt(time.Now().Unix(), 10)
	var fileName string
	fileName = fileNameOnly + timestampString + fileSuffix
	out, err := os.Create(basePath + fileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(basePath + fileName)
	}()
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	fileOut, err := file.Open()
	if err != nil {
		return err
	}
	_, err = io.Copy(out, fileOut)
	if err != nil {
		return err
	}

	f, err := excelize.OpenFile(basePath + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for i, row := range rows {
			if i == 0 {
				continue
			}
			loc, _ := time.LoadLocation("Asia/Shanghai")
			publishTimeStr := ""
			if row[2] != "" {
				publishTime, err := time.ParseInLocation("2006年1月2日", row[2], loc)
				if err != nil {
					return errors.New("请检查文件第" + strconv.FormatInt(int64(i+1), 10) + "行C列")
				}
				publishTimeStr = publishTime.Format("2006-01-02 15:04:05")
			}
			openTime, err := time.ParseInLocation("2006年1月2日 15:04:05", row[3]+" "+row[4], loc)
			if err != nil {
				return errors.New("请检查文件第" + strconv.FormatInt(int64(i+1), 10) + "行D和E列")
			}
			openTimeStr := openTime.Format("2006-01-02 15:04:05")
			endDateStr := ""
			if row[5] != "" {
				endDate, err := time.ParseInLocation("2006年1月2日", row[5], loc)
				if err != nil {
					return errors.New("请检查文件第" + strconv.FormatInt(int64(i+1), 10) + "行F列")
				}
				endDateStr = endDate.Format("2006-01-02 15:04:05")
			}
			projectDay, err := strconv.ParseInt(row[14], 10, 64)
			if err != nil {
				return errors.New("请检查文件第" + strconv.FormatInt(int64(i+1), 10) + "行O列")
			}
			projectAmountInRow, err := strconv.ParseFloat(row[15], 64)
			if err != nil {
				return errors.New("请检查文件第" + strconv.FormatInt(int64(i+1), 10) + "行P列")
			}
			projectAmount := projectAmountInRow * 10000
			tenderDepositInRow, err := strconv.ParseFloat(row[16], 64)
			if err != nil {
				return errors.New("请检查文件第" + strconv.FormatInt(int64(i+1), 10) + "行Q列")
			}
			tenderDeposit := tenderDepositInRow * 10000
			_, tendereeFile, err := f.GetCellHyperLink("Sheet1", fmt.Sprintf("R%d", i+1))
			isEnable := false
			isAutoMatic := false
			project := lg.Project{
				ProjectName:        &row[0],
				ProjectNo:          &row[1],
				ProjectAmount:      &projectAmount,
				TendereeName:       &row[10],
				TendereeAddress:    &row[11],
				TendereeTel:        &row[12],
				TendereeFile:       &tendereeFile,
				AgentTel:           &row[13],
				TenderDeposit:      &tenderDeposit,
				ProjectOpenTime:    &openTimeStr,
				ProjectPublishTime: &publishTimeStr,
				ProjectCity:        &row[6],
				ProjectCounty:      &row[7],
				ProjectDay:         &projectDay,
				TenderEndDate:      &endDateStr,
				ProjectType:        &row[8],
				ProjectCategory:    &row[9],
				IsEnable:           &isEnable,
				IsAutoMatic:        &isAutoMatic,
			}
			var p lg.Project
			if !errors.Is(tx.Where("project_no = ?", project.ProjectNo).First(&p).Error, gorm.ErrRecordNotFound) {
				continue
			}
			err = tx.Create(&project).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func AutoMaticUnEnableProject() {
	global.GVA_DB.Model(&lg.Project{}).Where("project_open_time is NOT NULL AND project_open_time != '' AND project_open_time < ?", time.Now().Format("2006-01-02 15:04:05")).Update("is_enable", false)
}
