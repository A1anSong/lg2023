package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
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
	if err != nil {
		return "", err
	}
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
