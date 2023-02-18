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

type RefundApi struct {
}

var refundService = service.ServiceGroupApp.LgServiceGroup.RefundService

func (refundApi *RefundApi) CreateRefund(c *gin.Context) {
	var refund lg.Refund
	err := c.ShouldBindJSON(&refund)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := refundService.CreateRefund(refund); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (refundApi *RefundApi) DeleteRefund(c *gin.Context) {
	var refund lg.Refund
	err := c.ShouldBindJSON(&refund)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := refundService.DeleteRefund(refund); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (refundApi *RefundApi) DeleteRefundByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := refundService.DeleteRefundByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (refundApi *RefundApi) UpdateRefund(c *gin.Context) {
	var refund lg.Refund
	err := c.ShouldBindJSON(&refund)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := refundService.UpdateRefund(refund); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (refundApi *RefundApi) FindRefund(c *gin.Context) {
	var refund lg.Refund
	err := c.ShouldBindQuery(&refund)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerefund, err := refundService.GetRefund(refund.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerefund": rerefund}, c)
	}
}

func (refundApi *RefundApi) GetRefundList(c *gin.Context) {
	var pageInfo lgReq.RefundSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := refundService.GetRefundInfoList(pageInfo); err != nil {
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
