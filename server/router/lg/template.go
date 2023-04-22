package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TemplateRouter struct {
}

func (s *TemplateRouter) InitTemplateRouter(Router *gin.RouterGroup) {
	templateRouter := Router.Group("template").Use(middleware.OperationRecord())
	templateRouterWithoutRecord := Router.Group("template")
	var templateApi = v1.ApiGroupApp.LgApiGroup.TemplateApi
	{
		templateRouter.POST("createTemplate", templateApi.CreateTemplate) // 新建Template
		templateRouter.PUT("updateTemplate", templateApi.UpdateTemplate)  // 更新Template
	}
	{
		templateRouterWithoutRecord.GET("getTemplateList", templateApi.GetTemplateList) // 获取Template列表
	}
}
