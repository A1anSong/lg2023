package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RefundRouter struct {
}

func (s *RefundRouter) InitRefundRouter(Router *gin.RouterGroup) {
	refundRouter := Router.Group("refund").Use(middleware.OperationRecord())
	refundRouterWithoutRecord := Router.Group("refund")
	var refundApi = v1.ApiGroupApp.LgApiGroup.RefundApi
	{
		refundRouter.POST("createRefund", refundApi.CreateRefund)             // 新建Refund
		refundRouter.DELETE("deleteRefund", refundApi.DeleteRefund)           // 删除Refund
		refundRouter.DELETE("deleteRefundByIds", refundApi.DeleteRefundByIds) // 批量删除Refund
		refundRouter.PUT("updateRefund", refundApi.UpdateRefund)              // 更新Refund
	}
	{
		refundRouterWithoutRecord.GET("findRefund", refundApi.FindRefund)       // 根据ID获取Refund
		refundRouterWithoutRecord.GET("getRefundList", refundApi.GetRefundList) // 获取Refund列表
	}
}
