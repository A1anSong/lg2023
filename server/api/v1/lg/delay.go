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

type DelayApi struct {
}

var delayService = service.ServiceGroupApp.LgServiceGroup.DelayService

func (delayApi *DelayApi) CreateDelay(c *gin.Context) {
	var delay lg.Delay
	err := c.ShouldBindJSON(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := delayService.CreateDelay(delay); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (delayApi *DelayApi) DeleteDelay(c *gin.Context) {
	var delay lg.Delay
	err := c.ShouldBindJSON(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := delayService.DeleteDelay(delay); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (delayApi *DelayApi) DeleteDelayByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := delayService.DeleteDelayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (delayApi *DelayApi) UpdateDelay(c *gin.Context) {
	var delay lg.Delay
	err := c.ShouldBindJSON(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := delayService.UpdateDelay(delay); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (delayApi *DelayApi) FindDelay(c *gin.Context) {
	var delay lg.Delay
	err := c.ShouldBindQuery(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redelay, err := delayService.GetDelay(delay.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redelay": redelay}, c)
	}
}

func (delayApi *DelayApi) GetDelayList(c *gin.Context) {
	var pageInfo lgReq.DelaySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := delayService.GetDelayInfoList(pageInfo); err != nil {
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
