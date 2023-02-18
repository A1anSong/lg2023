package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayRouter struct {
}

func (s *PayRouter) InitPayRouter(Router *gin.RouterGroup) {
	payRouter := Router.Group("pay").Use(middleware.OperationRecord())
	payRouterWithoutRecord := Router.Group("pay")
	var payApi = v1.ApiGroupApp.LgApiGroup.PayApi
	{
		payRouter.POST("createPay", payApi.CreatePay)             // 新建Pay
		payRouter.DELETE("deletePay", payApi.DeletePay)           // 删除Pay
		payRouter.DELETE("deletePayByIds", payApi.DeletePayByIds) // 批量删除Pay
		payRouter.PUT("updatePay", payApi.UpdatePay)              // 更新Pay
	}
	{
		payRouterWithoutRecord.GET("findPay", payApi.FindPay)       // 根据ID获取Pay
		payRouterWithoutRecord.GET("getPayList", payApi.GetPayList) // 获取Pay列表
	}
}
