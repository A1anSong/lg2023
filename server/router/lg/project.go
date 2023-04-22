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
		projectRouter.POST("createProject", projectApi.CreateProject)                // 新建Project
		projectRouter.PUT("updateProject", projectApi.UpdateProject)                 // 更新Project
		projectRouter.PUT("enableProjectByIds", projectApi.EnableProjectByIds)       // 批量上架Project
		projectRouter.PUT("autoMaticProjectByIds", projectApi.AutoMaticProjectByIds) // 批量自动审批Project
	}
	{
		projectRouterWithoutRecord.GET("findProject", projectApi.FindProject)                // 根据ID获取Project
		projectRouterWithoutRecord.GET("getProjectList", projectApi.GetProjectList)          // 获取Project列表
		projectRouterWithoutRecord.POST("bindProject", projectApi.BindProject)               // 绑定项目
		projectRouterWithoutRecord.POST("unbindProject", projectApi.UnbindProject)           // 解绑项目
		projectRouterWithoutRecord.POST("autoMaticProject", projectApi.AutoMaticProject)     // 设置项目自动化
		projectRouterWithoutRecord.POST("unAutoMaticProject", projectApi.UnAutoMaticProject) // 取消项目自动化
		projectRouterWithoutRecord.GET("downloadTemplate", projectApi.DownloadTemplate)      // 下载模板文件
		projectRouterWithoutRecord.POST("importExcel", projectApi.ImportExcel)               // 导入Excel
	}
}
