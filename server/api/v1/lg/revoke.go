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

type RevokeApi struct {
}

var revokeService = service.ServiceGroupApp.LgServiceGroup.RevokeService

func (revokeApi *RevokeApi) CreateRevoke(c *gin.Context) {
	var revoke lg.Revoke
	err := c.ShouldBindJSON(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := revokeService.CreateRevoke(revoke); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (revokeApi *RevokeApi) DeleteRevoke(c *gin.Context) {
	var revoke lg.Revoke
	err := c.ShouldBindJSON(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := revokeService.DeleteRevoke(revoke); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (revokeApi *RevokeApi) DeleteRevokeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := revokeService.DeleteRevokeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (revokeApi *RevokeApi) UpdateRevoke(c *gin.Context) {
	var revoke lg.Revoke
	err := c.ShouldBindJSON(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := revokeService.UpdateRevoke(revoke); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (revokeApi *RevokeApi) FindRevoke(c *gin.Context) {
	var revoke lg.Revoke
	err := c.ShouldBindQuery(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerevoke, err := revokeService.GetRevoke(revoke.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerevoke": rerevoke}, c)
	}
}

func (revokeApi *RevokeApi) GetRevokeList(c *gin.Context) {
	var pageInfo lgReq.RevokeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := revokeService.GetRevokeInfoList(pageInfo); err != nil {
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
