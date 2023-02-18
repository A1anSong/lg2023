package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type FileApi struct {
}

var fileService = service.ServiceGroupApp.LgServiceGroup.FileService

func (fileApi *FileApi) CreateFile(c *gin.Context) {
	var file lg.File
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileService.CreateFile(file); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (fileApi *FileApi) DeleteFile(c *gin.Context) {
	var file lg.File
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileService.DeleteFile(file); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (fileApi *FileApi) DeleteFileByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileService.DeleteFileByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (fileApi *FileApi) UpdateFile(c *gin.Context) {
	var file lg.File
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileService.UpdateFile(file); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (fileApi *FileApi) FindFile(c *gin.Context) {
	var file lg.File
	err := c.ShouldBindQuery(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refile, err := fileService.GetFile(file.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refile": refile}, c)
	}
}

func (fileApi *FileApi) GetFileList(c *gin.Context) {
	var pageInfo lgReq.FileSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileService.GetFileInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

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
