package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
)

type ClaimService struct {
}

func (claimService *ClaimService) CreateClaim(claim lg.Claim) (err error) {
	err = global.GVA_DB.Create(&claim).Error
	return err
}

func (claimService *ClaimService) DeleteClaim(claim lg.Claim) (err error) {
	err = global.GVA_DB.Delete(&claim).Error
	return err
}

func (claimService *ClaimService) DeleteClaimByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.Claim{}, "id in ?", ids.Ids).Error
	return err
}

func (claimService *ClaimService) UpdateClaim(claim lg.Claim) (err error) {
	err = global.GVA_DB.Save(&claim).Error
	return err
}

func (claimService *ClaimService) GetClaim(id uint) (claim lg.Claim, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&claim).Error
	return
}

func (claimService *ClaimService) GetClaimInfoList(info lgReq.ClaimSearch) (list []lg.Claim, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.Claim{})
	var claims []lg.Claim
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&claims).Error
	return claims, total, err
}
