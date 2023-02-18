package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RevokeRouter struct {
}

func (s *RevokeRouter) InitRevokeRouter(Router *gin.RouterGroup) {
	revokeRouter := Router.Group("revoke").Use(middleware.OperationRecord())
	revokeRouterWithoutRecord := Router.Group("revoke")
	var revokeApi = v1.ApiGroupApp.LgApiGroup.RevokeApi
	{
		revokeRouter.POST("createRevoke", revokeApi.CreateRevoke)             // 新建Revoke
		revokeRouter.DELETE("deleteRevoke", revokeApi.DeleteRevoke)           // 删除Revoke
		revokeRouter.DELETE("deleteRevokeByIds", revokeApi.DeleteRevokeByIds) // 批量删除Revoke
		revokeRouter.PUT("updateRevoke", revokeApi.UpdateRevoke)              // 更新Revoke
	}
	{
		revokeRouterWithoutRecord.GET("findRevoke", revokeApi.FindRevoke)       // 根据ID获取Revoke
		revokeRouterWithoutRecord.GET("getRevokeList", revokeApi.GetRevokeList) // 获取Revoke列表
	}
}
