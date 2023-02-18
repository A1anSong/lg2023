package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type FileService struct {
}

func (fileService *FileService) CreateFile(file lg.File) (err error) {
	err = global.GVA_DB.Create(&file).Error
	return err
}

func (fileService *FileService) DeleteFile(file lg.File) (err error) {
	err = global.GVA_DB.Delete(&file).Error
	return err
}

func (fileService *FileService) DeleteFileByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lg.File{}, "id in ?", ids.Ids).Error
	return err
}

func (fileService *FileService) UpdateFile(file lg.File) (err error) {
	err = global.GVA_DB.Save(&file).Error
	return err
}

func (fileService *FileService) GetFile(id uint) (file lg.File, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&file).Error
	return
}

func (fileService *FileService) GetFileInfoList(info lgReq.FileSearch) (list []lg.File, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lg.File{})
	var files []lg.File
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&files).Error
	return files, total, err
}

func (fileService *FileService) UploadFile(file *multipart.FileHeader) (fileName string, err error) {
	basePath := "./tmp/"
	fileSuffix := path.Ext(file.Filename)
	fileNameOnly := strings.TrimSuffix(file.Filename, fileSuffix)
	year, month, day := time.Now().Date()
	fileName = fileNameOnly + strconv.Itoa(year) + strconv.Itoa(int(month)) + strconv.Itoa(day) + fileSuffix
	out, err := os.Create(basePath + fileName)
	if err != nil {
		return "", err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	fileOut, err := file.Open()
	_, err = io.Copy(out, fileOut)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func (fileService *FileService) DownloadFile(id uint) (file lg.File, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&file).Error
	return
}
