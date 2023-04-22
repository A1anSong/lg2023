package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type InvoiceApplyRouter struct {
}

func (s *InvoiceApplyRouter) InitInvoiceApplyRouter(Router *gin.RouterGroup) {
	invoiceApplyRouterWithoutRecord := Router.Group("invoiceApply")
	var invoiceApplyApi = v1.ApiGroupApp.LgApiGroup.InvoiceApplyApi
	{
		invoiceApplyRouterWithoutRecord.GET("getInvoiceApplyList", invoiceApplyApi.GetInvoiceApplyList) // 获取InvoiceApply列表
	}
}
