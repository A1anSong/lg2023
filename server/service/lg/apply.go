package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
)

type ApplyService struct {
}

func (applyService *ApplyService) CreateApply(apply lg.Apply) (err error) {
	err = global.GVA_DB.Create(&apply).Error
	return err
}

func (applyService *ApplyService) DeleteApply(apply lg.Apply) (err error) {
	err = global.GVA_DB.Delete(&apply).Error
	return err
}

func (applyService *ApplyService) DeleteApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Apply{}, "id in ?", ids.Ids).Error
	return err
}

func (applyService *ApplyService) UpdateApply(apply lg.Apply) (err error) {
	err = global.GVA_DB.Save(&apply).Error
	return err
}

func (applyService *ApplyService) GetApply(id uint) (apply lg.Apply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&apply).Error
	return
}

func (applyService *ApplyService) GetApplyInfoList(info lgReq.ApplySearch) (list []lg.Apply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Apply{})
	var applys []lg.Apply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&applys).Error
	return applys, total, err
}
