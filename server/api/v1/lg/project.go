package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

type ProjectApi struct {
}

var projectService = service.ServiceGroupApp.LgServiceGroup.ProjectService

func (projectApi *ProjectApi) CreateProject(c *gin.Context) {
	var project lg.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := projectService.CreateProject(project); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (projectApi *ProjectApi) DeleteProject(c *gin.Context) {
	var project lg.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := projectService.DeleteProject(project); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (projectApi *ProjectApi) DeleteProjectByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := projectService.DeleteProjectByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (projectApi *ProjectApi) UpdateProject(c *gin.Context) {
	var project lg.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := projectService.UpdateProject(project); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (projectApi *ProjectApi) FindProject(c *gin.Context) {
	var project lg.Project
	err := c.ShouldBindQuery(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reproject, err := projectService.GetProject(project.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproject": reproject}, c)
	}
}

func (projectApi *ProjectApi) GetProjectList(c *gin.Context) {
	var pageInfo lgReq.ProjectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := projectService.GetProjectInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (projectApi *ProjectApi) BindProject(c *gin.Context) {
	var project lg.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := projectService.BindProject(project); err != nil {
		global.GVA_LOG.Error("绑定失败!", zap.Error(err))
		response.FailWithMessage("绑定失败："+err.Error(), c)
	} else {
		response.OkWithMessage("绑定成功", c)
	}
}

func (projectApi *ProjectApi) UnbindProject(c *gin.Context) {
	var project lg.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := projectService.UnbindProject(project); err != nil {
		global.GVA_LOG.Error("解绑失败!", zap.Error(err))
		response.FailWithMessage("解绑失败："+err.Error(), c)
	} else {
		response.OkWithMessage("解绑成功", c)
	}
}

func (projectApi *ProjectApi) AutoMaticProject(c *gin.Context) {
	var project lg.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := projectService.AutoMaticProject(project); err != nil {
		global.GVA_LOG.Error("绑定失败!", zap.Error(err))
		response.FailWithMessage("绑定失败："+err.Error(), c)
	} else {
		response.OkWithMessage("绑定成功", c)
	}
}

func (projectApi *ProjectApi) UnAutoMaticProject(c *gin.Context) {
	var project lg.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := projectService.UnAutoMaticProject(project); err != nil {
		global.GVA_LOG.Error("解绑失败!", zap.Error(err))
		response.FailWithMessage("解绑失败："+err.Error(), c)
	} else {
		response.OkWithMessage("解绑成功", c)
	}
}

func (projectApi *ProjectApi) DownloadTemplate(c *gin.Context) {
	fileName := c.Query("fileName")
	filePath := "./resource/project_template/" + fileName
	fi, err := os.Stat(filePath)
	if err != nil {
		global.GVA_LOG.Error("文件不存在!", zap.Error(err))
		response.FailWithMessage("文件不存在", c)
		return
	}
	if fi.IsDir() {
		global.GVA_LOG.Error("不支持下载文件夹!", zap.Error(err))
		response.FailWithMessage("不支持下载文件夹", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}

func (projectApi *ProjectApi) ImportExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	if err := projectService.ImportExcel(file); err != nil {
		global.GVA_LOG.Error("录入失败!"+err.Error(), zap.Error(err))
		response.FailWithMessage("录入失败："+err.Error(), c)
		return
	} else {
		response.OkWithMessage("录入成功", c)
	}
}
