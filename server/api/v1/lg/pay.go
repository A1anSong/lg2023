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

type PayApi struct {
}

var payService = service.ServiceGroupApp.LgServiceGroup.PayService

func (payApi *PayApi) CreatePay(c *gin.Context) {
	var pay lg.Pay
	err := c.ShouldBindJSON(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payService.CreatePay(pay); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (payApi *PayApi) DeletePay(c *gin.Context) {
	var pay lg.Pay
	err := c.ShouldBindJSON(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payService.DeletePay(pay); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (payApi *PayApi) DeletePayByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payService.DeletePayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (payApi *PayApi) UpdatePay(c *gin.Context) {
	var pay lg.Pay
	err := c.ShouldBindJSON(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payService.UpdatePay(pay); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (payApi *PayApi) FindPay(c *gin.Context) {
	var pay lg.Pay
	err := c.ShouldBindQuery(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repay, err := payService.GetPay(pay.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repay": repay}, c)
	}
}

func (payApi *PayApi) GetPayList(c *gin.Context) {
	var pageInfo lgReq.PaySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := payService.GetPayInfoList(pageInfo); err != nil {
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
