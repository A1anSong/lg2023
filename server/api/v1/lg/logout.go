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

type LogoutApi struct {
}

var logoutService = service.ServiceGroupApp.LgServiceGroup.LogoutService

func (logoutApi *LogoutApi) CreateLogout(c *gin.Context) {
	var logout lg.Logout
	err := c.ShouldBindJSON(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logoutService.CreateLogout(logout); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (logoutApi *LogoutApi) DeleteLogout(c *gin.Context) {
	var logout lg.Logout
	err := c.ShouldBindJSON(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logoutService.DeleteLogout(logout); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (logoutApi *LogoutApi) DeleteLogoutByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logoutService.DeleteLogoutByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (logoutApi *LogoutApi) UpdateLogout(c *gin.Context) {
	var logout lg.Logout
	err := c.ShouldBindJSON(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logoutService.UpdateLogout(logout); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (logoutApi *LogoutApi) FindLogout(c *gin.Context) {
	var logout lg.Logout
	err := c.ShouldBindQuery(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relogout, err := logoutService.GetLogout(logout.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relogout": relogout}, c)
	}
}

func (logoutApi *LogoutApi) GetLogoutList(c *gin.Context) {
	var pageInfo lgReq.LogoutSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := logoutService.GetLogoutInfoList(pageInfo); err != nil {
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
