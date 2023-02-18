package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LogoutRouter struct {
}

func (s *LogoutRouter) InitLogoutRouter(Router *gin.RouterGroup) {
	logoutRouter := Router.Group("logout").Use(middleware.OperationRecord())
	logoutRouterWithoutRecord := Router.Group("logout")
	var logoutApi = v1.ApiGroupApp.LgApiGroup.LogoutApi
	{
		logoutRouter.POST("createLogout", logoutApi.CreateLogout)             // 新建Logout
		logoutRouter.DELETE("deleteLogout", logoutApi.DeleteLogout)           // 删除Logout
		logoutRouter.DELETE("deleteLogoutByIds", logoutApi.DeleteLogoutByIds) // 批量删除Logout
		logoutRouter.PUT("updateLogout", logoutApi.UpdateLogout)              // 更新Logout
	}
	{
		logoutRouterWithoutRecord.GET("findLogout", logoutApi.FindLogout)       // 根据ID获取Logout
		logoutRouterWithoutRecord.GET("getLogoutList", logoutApi.GetLogoutList) // 获取Logout列表
	}
}
