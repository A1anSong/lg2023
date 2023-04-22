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
		orderRouter.PUT("approveApply", orderApi.ApproveApply) // 审批同意申请
		orderRouter.PUT("rejectApply", orderApi.RejectApply)   // 审批拒绝申请
		orderRouter.PUT("approveDelay", orderApi.ApproveDelay) // 审批同意延期
		orderRouter.PUT("rejectDelay", orderApi.RejectDelay)   // 审批拒绝延期
		orderRouter.PUT("approveClaim", orderApi.ApproveClaim) // 审批同意理赔
		orderRouter.PUT("rejectClaim", orderApi.RejectClaim)   // 审批拒绝理赔
		orderRouter.PUT("openLetter", orderApi.OpenLetter)     // 提交开函申请
		orderRouter.PUT("rePush", orderApi.RePush)             // 提交重推申请
		orderRouter.PUT("assignOrder", orderApi.AssignOrder)   // 订单认领
	}
	{
		orderRouterWithoutRecord.GET("getOrderList", orderApi.GetOrderList)                         // 获取Order列表
		orderRouterWithoutRecord.GET("getInsuranceBalance", orderApi.GetInsuranceBalance)           // 获取担保额度
		orderRouterWithoutRecord.GET("getOrderStatisticData", orderApi.GetOrderStatisticData)       // 获取Order统计数据
		orderRouterWithoutRecord.GET("getEmployeeStatisticData", orderApi.GetEmployeeStatisticData) // 获取Order开单数据
		orderRouterWithoutRecord.GET("getGEODistributionData", orderApi.GetGEODistributionData)     // 获取Order地理分布数据
		orderRouterWithoutRecord.GET("getOrderTrendData", orderApi.GetOrderTrendData)               // 获取Order趋势数据
		orderRouterWithoutRecord.GET("exportExcel", orderApi.ExportExcel)                           // 导出Order数据到excel
		orderRouterWithoutRecord.POST("markOfflineRefund", orderApi.MarkOfflineRefund)              // 绑定项目
		orderRouterWithoutRecord.POST("unmarkOfflineRefund", orderApi.UnmarkOfflineRefund)          // 解绑项目
		orderRouterWithoutRecord.GET("exportInvoiceExcel", orderApi.ExportInvoiceExcel)             // 导出Invoice数据到excel
	}
}

func (s *OrderRouter) InitPublicOrderRouter(Router *gin.RouterGroup) {
	orderRouterWithoutRecord := Router.Group("order")
	var orderApi = v1.ApiGroupApp.LgApiGroup.OrderApi
	orderRouterWithoutRecord.POST("elogValidate", orderApi.ElogValidate) // 保函鉴真
}
