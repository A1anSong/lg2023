package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type FileRouter struct {
}

func (s *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	fileRouterWithoutRecord := Router.Group("file")
	var fileApi = v1.ApiGroupApp.LgApiGroup.FileApi
	{
		fileRouterWithoutRecord.POST("upload", fileApi.UploadFile)    // 上传文件
		fileRouterWithoutRecord.GET("download", fileApi.DownloadFile) // 下载文件
	}
}
