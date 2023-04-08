package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

func (s *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	orderRouter := Router.Group("order").Use(middleware.OperationRecord())
	orderRouterWithoutRecord := Router.Group("order")
	var orderApi = v1.ApiGroupApp.LgApiGroup.OrderApi
	{
		orderRouter.POST("createOrder", orderApi.CreateOrder)             // 新建Order
		orderRouter.DELETE("deleteOrder", orderApi.DeleteOrder)           // 删除Order
		orderRouter.DELETE("deleteOrderByIds", orderApi.DeleteOrderByIds) // 批量删除Order
		orderRouter.PUT("approveApply", orderApi.ApproveApply)            // 审批同意申请
		orderRouter.PUT("rejectApply", orderApi.RejectApply)              // 审批拒绝申请
		orderRouter.PUT("approveDelay", orderApi.ApproveDelay)            // 审批同意延期
		orderRouter.PUT("rejectDelay", orderApi.RejectDelay)              // 审批拒绝延期
		orderRouter.PUT("approveRefund", orderApi.ApproveRefund)          // 审批同意退函
		orderRouter.PUT("rejectRefund", orderApi.RejectRefund)            // 审批拒绝退函
		orderRouter.PUT("approveClaim", orderApi.ApproveClaim)            // 审批同意理赔
		orderRouter.PUT("rejectClaim", orderApi.RejectClaim)              // 审批拒绝理赔
		orderRouter.PUT("openLetter", orderApi.OpenLetter)                // 提交开函申请
		orderRouter.PUT("rePush", orderApi.RePush)                        // 提交重推申请
		orderRouter.PUT("requestInvoice", orderApi.RequestInvoice)        // 提交发票申请
		orderRouter.PUT("queryInvoice", orderApi.QueryInvoice)            // 查询开票结果
		orderRouter.PUT("assignOrder", orderApi.AssignOrder)              // 订单认领
	}
	{
		orderRouterWithoutRecord.GET("findOrder", orderApi.FindOrder)                         // 根据ID获取Order
		orderRouterWithoutRecord.GET("getOrderList", orderApi.GetOrderList)                   // 获取Order列表
		orderRouterWithoutRecord.GET("getOrderStatisticData", orderApi.GetOrderStatisticData) // 获取Order统计数据
		orderRouterWithoutRecord.GET("exportExcel", orderApi.ExportExcel)                     // 导出Order数据到excel
		orderRouterWithoutRecord.GET("findOrderByNos", orderApi.FindOrderByNos)               // 根据OrderNo获取Order
		orderRouterWithoutRecord.POST("markOfflineRefund", orderApi.MarkOfflineRefund)        // 绑定项目
		orderRouterWithoutRecord.POST("unmarkOfflineRefund", orderApi.UnmarkOfflineRefund)    // 解绑项目
		orderRouterWithoutRecord.GET("exportInvoiceExcel", orderApi.ExportInvoiceExcel)       // 导出Invoice数据到excel
	}
}
