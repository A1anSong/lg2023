package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
)

type RevokeService struct {
}

func (revokeService *RevokeService) CreateRevoke(revoke lg.Revoke) (err error) {
	err = global.GVA_DB.Create(&revoke).Error
	return err
}

func (revokeService *RevokeService) DeleteRevoke(revoke lg.Revoke) (err error) {
	err = global.GVA_DB.Delete(&revoke).Error
	return err
}

func (revokeService *RevokeService) DeleteRevokeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Revoke{}, "id in ?", ids.Ids).Error
	return err
}

func (revokeService *RevokeService) UpdateRevoke(revoke lg.Revoke) (err error) {
	err = global.GVA_DB.Save(&revoke).Error
	return err
}

func (revokeService *RevokeService) GetRevoke(id uint) (revoke lg.Revoke, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&revoke).Error
	return
}

func (revokeService *RevokeService) GetRevokeInfoList(info lgReq.RevokeSearch) (list []lg.Revoke, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Revoke{})
	var revokes []lg.Revoke
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&revokes).Error
	return revokes, total, err
}
