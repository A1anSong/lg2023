package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
)

type DelayService struct {
}

func (delayService *DelayService) CreateDelay(delay lg.Delay) (err error) {
	err = global.GVA_DB.Create(&delay).Error
	return err
}

func (delayService *DelayService) DeleteDelay(delay lg.Delay) (err error) {
	err = global.GVA_DB.Delete(&delay).Error
	return err
}

func (delayService *DelayService) DeleteDelayByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Delay{}, "id in ?", ids.Ids).Error
	return err
}

func (delayService *DelayService) UpdateDelay(delay lg.Delay) (err error) {
	err = global.GVA_DB.Save(&delay).Error
	return err
}

func (delayService *DelayService) GetDelay(id uint) (delay lg.Delay, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&delay).Error
	return
}

func (delayService *DelayService) GetDelayInfoList(info lgReq.DelaySearch) (list []lg.Delay, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Delay{})
	var delays []lg.Delay
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&delays).Error
	return delays, total, err
}
