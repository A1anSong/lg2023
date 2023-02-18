package lg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LetterRouter struct {
}

func (s *LetterRouter) InitLetterRouter(Router *gin.RouterGroup) {
	letterRouter := Router.Group("letter").Use(middleware.OperationRecord())
	letterRouterWithoutRecord := Router.Group("letter")
	var letterApi = v1.ApiGroupApp.LgApiGroup.LetterApi
	{
		letterRouter.POST("createLetter", letterApi.CreateLetter)             // 新建Letter
		letterRouter.DELETE("deleteLetter", letterApi.DeleteLetter)           // 删除Letter
		letterRouter.DELETE("deleteLetterByIds", letterApi.DeleteLetterByIds) // 批量删除Letter
		letterRouter.PUT("updateLetter", letterApi.UpdateLetter)              // 更新Letter
	}
	{
		letterRouterWithoutRecord.GET("findLetter", letterApi.FindLetter)       // 根据ID获取Letter
		letterRouterWithoutRecord.GET("getLetterList", letterApi.GetLetterList) // 获取Letter列表
	}
}
