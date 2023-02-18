package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
)

type LetterService struct {
}

func (letterService *LetterService) CreateLetter(letter lg.Letter) (err error) {
	err = global.GVA_DB.Create(&letter).Error
	return err
}

func (letterService *LetterService) DeleteLetter(letter lg.Letter) (err error) {
	err = global.GVA_DB.Delete(&letter).Error
	return err
}

func (letterService *LetterService) DeleteLetterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Letter{}, "id in ?", ids.Ids).Error
	return err
}

func (letterService *LetterService) UpdateLetter(letter lg.Letter) (err error) {
	err = global.GVA_DB.Save(&letter).Error
	return err
}

func (letterService *LetterService) GetLetter(id uint) (letter lg.Letter, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&letter).Error
	return
}

func (letterService *LetterService) GetLetterInfoList(info lgReq.LetterSearch) (list []lg.Letter, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Letter{})
	var letters []lg.Letter
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&letters).Error
	return letters, total, err
}
