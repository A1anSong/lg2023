package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InvoiceApplyRouter struct {
}

func (s *InvoiceApplyRouter) InitInvoiceApplyRouter(Router *gin.RouterGroup) {
	invoiceApplyRouter := Router.Group("invoiceApply").Use(middleware.OperationRecord())
	invoiceApplyRouterWithoutRecord := Router.Group("invoiceApply")
	var invoiceApplyApi = v1.ApiGroupApp.LgApiGroup.InvoiceApplyApi
	{
		invoiceApplyRouter.POST("createInvoiceApply", invoiceApplyApi.CreateInvoiceApply)             // 新建InvoiceApply
		invoiceApplyRouter.DELETE("deleteInvoiceApply", invoiceApplyApi.DeleteInvoiceApply)           // 删除InvoiceApply
		invoiceApplyRouter.DELETE("deleteInvoiceApplyByIds", invoiceApplyApi.DeleteInvoiceApplyByIds) // 批量删除InvoiceApply
		invoiceApplyRouter.PUT("updateInvoiceApply", invoiceApplyApi.UpdateInvoiceApply)              // 更新InvoiceApply
		invoiceApplyRouter.PUT("approveInvoiceApply", invoiceApplyApi.ApproveInvoiceApply)            // 审批同意发票申请
		invoiceApplyRouter.PUT("rejectInvoiceApply", invoiceApplyApi.RejectInvoiceApply)              // 审批拒绝发票申请
	}
	{
		invoiceApplyRouterWithoutRecord.GET("findInvoiceApply", invoiceApplyApi.FindInvoiceApply)       // 根据ID获取InvoiceApply
		invoiceApplyRouterWithoutRecord.GET("getInvoiceApplyList", invoiceApplyApi.GetInvoiceApplyList) // 获取InvoiceApply列表
	}
}
