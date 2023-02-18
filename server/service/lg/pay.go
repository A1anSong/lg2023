package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
)

type PayService struct {
}

func (payService *PayService) CreatePay(pay lg.Pay) (err error) {
	err = global.GVA_DB.Create(&pay).Error
	return err
}

func (payService *PayService) DeletePay(pay lg.Pay) (err error) {
	err = global.GVA_DB.Delete(&pay).Error
	return err
}

func (payService *PayService) DeletePayByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Pay{}, "id in ?", ids.Ids).Error
	return err
}

func (payService *PayService) UpdatePay(pay lg.Pay) (err error) {
	err = global.GVA_DB.Save(&pay).Error
	return err
}

func (payService *PayService) GetPay(id uint) (pay lg.Pay, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&pay).Error
	return
}

func (payService *PayService) GetPayInfoList(info lgReq.PaySearch) (list []lg.Pay, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Pay{})
	var pays []lg.Pay
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&pays).Error
	return pays, total, err
}
