package middleware

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/lg"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func JRValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 联调后删除
		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%v\n", string(data))
		//很关键
		//把读过的字节流重新放到body
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

		var request jrrequest.JRRequest
		err = c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingParam),
				Msg:  jrapi.MissingParam.String(),
				Data: "",
			})
			c.Abort()
			return
		}
		// sm3验签
		if request.Signature != lg.SM3Sign("appKey="+request.AppKey+"&data="+request.Data+"&requestId="+request.RequestId+"&timestamp="+request.Timestamp) {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.SignCheckFailed),
				Msg:  jrapi.SignCheckFailed.String(),
				Data: "",
			})
			c.Abort()
			return
		}
		// 提取data
		byteEncryptData, err := base64.StdEncoding.DecodeString(request.Data)
		if err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingParam),
				Msg:  jrapi.MissingParam.String(),
				Data: "",
			})
			c.Abort()
			return
		}
		// sm4解密
		jsonData, err := lg.Sm4Decrypt(byteEncryptData)
		if err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.DecryptFailed),
				Msg:  jrapi.DecryptFailed.String(),
				Data: "",
			})
			c.Abort()
			return
		}
		fmt.Println(jsonData)
		c.Set("jsonData", jsonData)
		c.Next()
	}
}
