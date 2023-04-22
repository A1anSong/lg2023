package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type FileApi struct {
}

var fileService = service.ServiceGroupApp.LgServiceGroup.FileService

func (fileApi *FileApi) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	if fileName, err := fileService.UploadFile(file); err != nil {
		global.GVA_LOG.Error("写入文件失败!", zap.Error(err))
		response.FailWithMessage("写入文件失败", c)
		return
	} else {
		response.OkWithDetailed(gin.H{
			"fileName": fileName,
		}, "获取成功", c)
	}
}

func (fileApi *FileApi) DownloadFile(c *gin.Context) {
	var file lg.File
	err := c.ShouldBindQuery(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refile, err := fileService.DownloadFile(file.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		c.Writer.Header().Add("success", "true")
		c.Header("Content-Disposition", "attachment; filename="+*refile.FileName) // 用来指定下载下来的文件名
		c.Data(http.StatusOK, "application/octet-stream", refile.FileSteam)
	}
}
