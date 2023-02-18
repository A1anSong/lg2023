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

type InvoiceApi struct {
}

var invoiceService = service.ServiceGroupApp.LgServiceGroup.InvoiceService

func (invoiceApi *InvoiceApi) CreateInvoice(c *gin.Context) {
	var invoice lg.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceService.CreateInvoice(invoice); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (invoiceApi *InvoiceApi) DeleteInvoice(c *gin.Context) {
	var invoice lg.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceService.DeleteInvoice(invoice); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (invoiceApi *InvoiceApi) DeleteInvoiceByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceService.DeleteInvoiceByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (invoiceApi *InvoiceApi) UpdateInvoice(c *gin.Context) {
	var invoice lg.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceService.UpdateInvoice(invoice); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (invoiceApi *InvoiceApi) FindInvoice(c *gin.Context) {
	var invoice lg.Invoice
	err := c.ShouldBindQuery(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reinvoice, err := invoiceService.GetInvoice(invoice.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinvoice": reinvoice}, c)
	}
}

func (invoiceApi *InvoiceApi) GetInvoiceList(c *gin.Context) {
	var pageInfo lgReq.InvoiceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := invoiceService.GetInvoiceInfoList(pageInfo); err != nil {
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
