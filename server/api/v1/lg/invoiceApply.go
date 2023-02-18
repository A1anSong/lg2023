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

type InvoiceApplyApi struct {
}

var invoiceApplyService = service.ServiceGroupApp.LgServiceGroup.InvoiceApplyService

func (invoiceApplyApi *InvoiceApplyApi) CreateInvoiceApply(c *gin.Context) {
	var invoiceApply lg.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceApplyService.CreateInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (invoiceApplyApi *InvoiceApplyApi) DeleteInvoiceApply(c *gin.Context) {
	var invoiceApply lg.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceApplyService.DeleteInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (invoiceApplyApi *InvoiceApplyApi) DeleteInvoiceApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceApplyService.DeleteInvoiceApplyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (invoiceApplyApi *InvoiceApplyApi) UpdateInvoiceApply(c *gin.Context) {
	var invoiceApply lg.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceApplyService.UpdateInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (invoiceApplyApi *InvoiceApplyApi) FindInvoiceApply(c *gin.Context) {
	var invoiceApply lg.InvoiceApply
	err := c.ShouldBindQuery(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reinvoiceApply, err := invoiceApplyService.GetInvoiceApply(invoiceApply.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinvoiceApply": reinvoiceApply}, c)
	}
}

func (invoiceApplyApi *InvoiceApplyApi) GetInvoiceApplyList(c *gin.Context) {
	var pageInfo lgReq.InvoiceApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := invoiceApplyService.GetInvoiceApplyInfoList(pageInfo); err != nil {
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

func (invoiceApplyApi *InvoiceApplyApi) ApproveInvoiceApply(c *gin.Context) {
	var invoiceApply lg.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceApplyService.ApproveInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (invoiceApplyApi *InvoiceApplyApi) RejectInvoiceApply(c *gin.Context) {
	var invoiceApply lg.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceApplyService.RejectInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
