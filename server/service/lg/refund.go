package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
)

type RefundService struct {
}

func (refundService *RefundService) CreateRefund(refund lg.Refund) (err error) {
	err = global.GVA_DB.Create(&refund).Error
	return err
}

func (refundService *RefundService) DeleteRefund(refund lg.Refund) (err error) {
	err = global.GVA_DB.Delete(&refund).Error
	return err
}

func (refundService *RefundService) DeleteRefundByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Refund{}, "id in ?", ids.Ids).Error
	return err
}

func (refundService *RefundService) UpdateRefund(refund lg.Refund) (err error) {
	err = global.GVA_DB.Save(&refund).Error
	return err
}

func (refundService *RefundService) GetRefund(id uint) (refund lg.Refund, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&refund).Error
	return
}

func (refundService *RefundService) GetRefundInfoList(info lgReq.RefundSearch) (list []lg.Refund, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Refund{})
	var refunds []lg.Refund
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&refunds).Error
	return refunds, total, err
}
