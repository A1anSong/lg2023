package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type JRAPIRouter struct {
}

func (s *JRAPIRouter) InitJRAPIRouter(Router *gin.RouterGroup) {
	jRAPIRouter := Router.Group("lg").Use(middleware.JRValidate())
	jRAPIRouterPublic := Router.Group("lg")
	jRApi := v1.ApiGroupApp.JRAPI
	{
		jRAPIRouter.POST("apply", jRApi.Apply)
		jRAPIRouter.POST("payPush", jRApi.PayPush)
		jRAPIRouter.POST("queryInfo", jRApi.QueryInfo)
		jRAPIRouter.POST("revokePush", jRApi.RevokePush)
		jRAPIRouter.POST("applyDelay", jRApi.ApplyDelay)
		jRAPIRouter.POST("applyRefund", jRApi.ApplyRefund)
		jRAPIRouter.POST("applyClaim", jRApi.ApplyClaim)
		jRAPIRouter.POST("logoutPush", jRApi.LogoutPush)
		jRAPIRouter.POST("applyInvoice", jRApi.ApplyInvoice)
	}
	{
		jRAPIRouterPublic.GET("letterFileDownload", jRApi.LetterFileDownload)
		jRAPIRouterPublic.GET("delayFileDownload", jRApi.DelayFileDownload)
		jRAPIRouterPublic.GET("invoiceFileDownload", jRApi.InvoiceFileDownload)
	}
}
