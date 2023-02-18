package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApplyRouter struct {
}

func (s *ApplyRouter) InitApplyRouter(Router *gin.RouterGroup) {
	applyRouter := Router.Group("apply").Use(middleware.OperationRecord())
	applyRouterWithoutRecord := Router.Group("apply")
	var applyApi = v1.ApiGroupApp.LgApiGroup.ApplyApi
	{
		applyRouter.POST("createApply", applyApi.CreateApply)             // 新建Apply
		applyRouter.DELETE("deleteApply", applyApi.DeleteApply)           // 删除Apply
		applyRouter.DELETE("deleteApplyByIds", applyApi.DeleteApplyByIds) // 批量删除Apply
		applyRouter.PUT("updateApply", applyApi.UpdateApply)              // 更新Apply
	}
	{
		applyRouterWithoutRecord.GET("findApply", applyApi.FindApply)       // 根据ID获取Apply
		applyRouterWithoutRecord.GET("getApplyList", applyApi.GetApplyList) // 获取Apply列表
	}
}
