package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClaimRouter struct {
}

func (s *ClaimRouter) InitClaimRouter(Router *gin.RouterGroup) {
	claimRouter := Router.Group("claim").Use(middleware.OperationRecord())
	claimRouterWithoutRecord := Router.Group("claim")
	var claimApi = v1.ApiGroupApp.LgApiGroup.ClaimApi
	{
		claimRouter.POST("createClaim", claimApi.CreateClaim)             // 新建Claim
		claimRouter.DELETE("deleteClaim", claimApi.DeleteClaim)           // 删除Claim
		claimRouter.DELETE("deleteClaimByIds", claimApi.DeleteClaimByIds) // 批量删除Claim
		claimRouter.PUT("updateClaim", claimApi.UpdateClaim)              // 更新Claim
	}
	{
		claimRouterWithoutRecord.GET("findClaim", claimApi.FindClaim)       // 根据ID获取Claim
		claimRouterWithoutRecord.GET("getClaimList", claimApi.GetClaimList) // 获取Claim列表
	}
}
