package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DelayRouter struct {
}

func (s *DelayRouter) InitDelayRouter(Router *gin.RouterGroup) {
	delayRouter := Router.Group("delay").Use(middleware.OperationRecord())
	delayRouterWithoutRecord := Router.Group("delay")
	var delayApi = v1.ApiGroupApp.LgApiGroup.DelayApi
	{
		delayRouter.POST("createDelay", delayApi.CreateDelay)             // 新建Delay
		delayRouter.DELETE("deleteDelay", delayApi.DeleteDelay)           // 删除Delay
		delayRouter.DELETE("deleteDelayByIds", delayApi.DeleteDelayByIds) // 批量删除Delay
		delayRouter.PUT("updateDelay", delayApi.UpdateDelay)              // 更新Delay
	}
	{
		delayRouterWithoutRecord.GET("findDelay", delayApi.FindDelay)       // 根据ID获取Delay
		delayRouterWithoutRecord.GET("getDelayList", delayApi.GetDelayList) // 获取Delay列表
	}
}
