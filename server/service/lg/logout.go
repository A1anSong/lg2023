package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
)

type LogoutService struct {
}

func (logoutService *LogoutService) CreateLogout(logout lg.Logout) (err error) {
	err = global.GVA_DB.Create(&logout).Error
	return err
}

func (logoutService *LogoutService) DeleteLogout(logout lg.Logout) (err error) {
	err = global.GVA_DB.Delete(&logout).Error
	return err
}

func (logoutService *LogoutService) DeleteLogoutByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Logout{}, "id in ?", ids.Ids).Error
	return err
}

func (logoutService *LogoutService) UpdateLogout(logout lg.Logout) (err error) {
	err = global.GVA_DB.Save(&logout).Error
	return err
}

func (logoutService *LogoutService) GetLogout(id uint) (logout lg.Logout, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&logout).Error
	return
}

func (logoutService *LogoutService) GetLogoutInfoList(info lgReq.LogoutSearch) (list []lg.Logout, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Logout{})
	var logouts []lg.Logout
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&logouts).Error
	return logouts, total, err
}
