package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectRouter struct {
}

func (s *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) {
	projectRouter := Router.Group("project").Use(middleware.OperationRecord())
	projectRouterWithoutRecord := Router.Group("project")
	var projectApi = v1.ApiGroupApp.LgApiGroup.ProjectApi
	{
		projectRouter.POST("createProject", projectApi.CreateProject)             // 新建Project
		projectRouter.DELETE("deleteProject", projectApi.DeleteProject)           // 删除Project
		projectRouter.DELETE("deleteProjectByIds", projectApi.DeleteProjectByIds) // 批量删除Project
		projectRouter.PUT("updateProject", projectApi.UpdateProject)              // 更新Project
	}
	{
		projectRouterWithoutRecord.GET("findProject", projectApi.FindProject)       // 根据ID获取Project
		projectRouterWithoutRecord.GET("getProjectList", projectApi.GetProjectList) // 获取Project列表
		projectRouterWithoutRecord.POST("bindProject", projectApi.BindProject)      // 绑定项目
		projectRouterWithoutRecord.POST("unbindProject", projectApi.UnbindProject)  // 解绑项目
	}
}
