package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InvoiceRouter struct {
}

func (s *InvoiceRouter) InitInvoiceRouter(Router *gin.RouterGroup) {
	invoiceRouter := Router.Group("invoice").Use(middleware.OperationRecord())
	invoiceRouterWithoutRecord := Router.Group("invoice")
	var invoiceApi = v1.ApiGroupApp.LgApiGroup.InvoiceApi
	{
		invoiceRouter.POST("createInvoice", invoiceApi.CreateInvoice)             // 新建Invoice
		invoiceRouter.DELETE("deleteInvoice", invoiceApi.DeleteInvoice)           // 删除Invoice
		invoiceRouter.DELETE("deleteInvoiceByIds", invoiceApi.DeleteInvoiceByIds) // 批量删除Invoice
		invoiceRouter.PUT("updateInvoice", invoiceApi.UpdateInvoice)              // 更新Invoice
	}
	{
		invoiceRouterWithoutRecord.GET("findInvoice", invoiceApi.FindInvoice)       // 根据ID获取Invoice
		invoiceRouterWithoutRecord.GET("getInvoiceList", invoiceApi.GetInvoiceList) // 获取Invoice列表
	}
}
