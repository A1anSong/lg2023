package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/nonmigrate"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TemplateApi struct {
}

var templateService = service.ServiceGroupApp.LgServiceGroup.TemplateService

func (templateApi *TemplateApi) CreateTemplate(c *gin.Context) {
	var templateAndFile nonmigrate.TemplateAndFile
	err := c.ShouldBindJSON(&templateAndFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.CreateTemplate(templateAndFile); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (templateApi *TemplateApi) DeleteTemplate(c *gin.Context) {
	var template lg.Template
	err := c.ShouldBindJSON(&template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.DeleteTemplate(template); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (templateApi *TemplateApi) DeleteTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.DeleteTemplateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (templateApi *TemplateApi) UpdateTemplate(c *gin.Context) {
	var template lg.Template
	err := c.ShouldBindJSON(&template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.UpdateTemplate(template); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (templateApi *TemplateApi) FindTemplate(c *gin.Context) {
	var template lg.Template
	err := c.ShouldBindQuery(&template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retemplate, err := templateService.GetTemplate(template.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retemplate": retemplate}, c)
	}
}

func (templateApi *TemplateApi) GetTemplateList(c *gin.Context) {
	var pageInfo lgReq.TemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := templateService.GetTemplateInfoList(pageInfo); err != nil {
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
