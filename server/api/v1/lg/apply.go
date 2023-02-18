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
)

type ApplyApi struct {
}

var applyService = service.ServiceGroupApp.LgServiceGroup.ApplyService

func (applyApi *ApplyApi) CreateApply(c *gin.Context) {
	var apply lg.Apply
	err := c.ShouldBindJSON(&apply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := applyService.CreateApply(apply); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (applyApi *ApplyApi) DeleteApply(c *gin.Context) {
	var apply lg.Apply
	err := c.ShouldBindJSON(&apply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := applyService.DeleteApply(apply); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (applyApi *ApplyApi) DeleteApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := applyService.DeleteApplyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (applyApi *ApplyApi) UpdateApply(c *gin.Context) {
	var apply lg.Apply
	err := c.ShouldBindJSON(&apply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := applyService.UpdateApply(apply); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (applyApi *ApplyApi) FindApply(c *gin.Context) {
	var apply lg.Apply
	err := c.ShouldBindQuery(&apply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reapply, err := applyService.GetApply(apply.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapply": reapply}, c)
	}
}

func (applyApi *ApplyApi) GetApplyList(c *gin.Context) {
	var pageInfo lgReq.ApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := applyService.GetApplyInfoList(pageInfo); err != nil {
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
