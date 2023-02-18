package lg

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/lg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type JRAPI struct {
}

var jrAPIService = service.ServiceGroupApp.JRAPIService

func (jrAPI *JRAPI) Apply(c *gin.Context) {
	var reApply jrrequest.JRAPIApply
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reApply)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resApply jrresponse.JRAPIApply
	if resApply, err = jrAPIService.ApplyOrder(reApply); err != nil {
		global.GVA_LOG.Error("创建Apply失败!", zap.Error(err))
		if err.Error() == jrapi.MissingServiceParam.String() {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingServiceParam),
				Msg:  jrapi.MissingServiceParam.String(),
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.FAILED),
				Msg:  err.Error(),
				Data: "",
			})
		}
	} else {
		var response jrresponse.JRResponse
		if response, err = lg.GenJRResponse(resApply); err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: -1,
				Msg:  "resApply序列化失败",
				Data: "",
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func (jrAPI *JRAPI) PayPush(c *gin.Context) {
	var rePayPush jrrequest.JRAPIPayPush
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &rePayPush)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resPayPush jrresponse.JRAPIPayPush
	if resPayPush, err = jrAPIService.PayPush(rePayPush); err != nil {
		global.GVA_LOG.Error("创建Pay失败!", zap.Error(err))
		if err.Error() == jrapi.MissingServiceParam.String() {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingServiceParam),
				Msg:  jrapi.MissingServiceParam.String(),
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.FAILED),
				Msg:  err.Error(),
				Data: "",
			})
		}
	} else {
		var response jrresponse.JRResponse
		if response, err = lg.GenJRResponse(resPayPush); err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: -1,
				Msg:  "resApply序列化失败",
				Data: "",
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func (jrAPI *JRAPI) QueryInfo(c *gin.Context) {
	var reQueryInfo jrrequest.JRAPIQueryInfo
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reQueryInfo)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resQueryInfo jrresponse.JRAPIQueryInfo
	if resQueryInfo, err = jrAPIService.QueryInfo(reQueryInfo); err != nil {
		global.GVA_LOG.Error("查询QueryInfo失败!", zap.Error(err))
		if err.Error() == jrapi.MissingServiceParam.String() {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingServiceParam),
				Msg:  jrapi.MissingServiceParam.String(),
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.FAILED),
				Msg:  err.Error(),
				Data: "",
			})
		}
	} else {
		var response jrresponse.JRResponse
		if response, err = lg.GenJRResponse(resQueryInfo); err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: -1,
				Msg:  "resApply序列化失败",
				Data: "",
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func (jrAPI *JRAPI) RevokePush(c *gin.Context) {
	var reRevokePush jrrequest.JRAPIRevokePush
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reRevokePush)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resRevokePush jrresponse.JRAPIRevokePush
	if resRevokePush, err = jrAPIService.RevokePush(reRevokePush); err != nil {
		global.GVA_LOG.Error("创建Revoke失败!", zap.Error(err))
		if err.Error() == jrapi.MissingServiceParam.String() {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingServiceParam),
				Msg:  jrapi.MissingServiceParam.String(),
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.FAILED),
				Msg:  err.Error(),
				Data: "",
			})
		}
	} else {
		var response jrresponse.JRResponse
		if response, err = lg.GenJRResponse(resRevokePush); err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: -1,
				Msg:  "resApply序列化失败",
				Data: "",
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func (jrAPI *JRAPI) ApplyDelay(c *gin.Context) {
	var reApplyDelay jrrequest.JRAPIApplyDelay
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reApplyDelay)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resApplyDelay jrresponse.JRAPIApplyDelay
	if resApplyDelay, err = jrAPIService.ApplyDelay(reApplyDelay); err != nil {
		global.GVA_LOG.Error("创建Delay失败!", zap.Error(err))
		if err.Error() == jrapi.MissingServiceParam.String() {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingServiceParam),
				Msg:  jrapi.MissingServiceParam.String(),
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.FAILED),
				Msg:  err.Error(),
				Data: "",
			})
		}
	} else {
		var response jrresponse.JRResponse
		if response, err = lg.GenJRResponse(resApplyDelay); err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: -1,
				Msg:  "resApply序列化失败",
				Data: "",
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func (jrAPI *JRAPI) ApplyRefund(c *gin.Context) {
	var reApplyRefund jrrequest.JRAPIApplyRefund
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reApplyRefund)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resApplyRefund jrresponse.JRAPIApplyRefund
	if resApplyRefund, err = jrAPIService.ApplyRefund(reApplyRefund); err != nil {
		global.GVA_LOG.Error("创建Refund失败!", zap.Error(err))
		if err.Error() == jrapi.MissingServiceParam.String() {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingServiceParam),
				Msg:  jrapi.MissingServiceParam.String(),
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.FAILED),
				Msg:  err.Error(),
				Data: "",
			})
		}
	} else {
		var response jrresponse.JRResponse
		if response, err = lg.GenJRResponse(resApplyRefund); err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: -1,
				Msg:  "resApply序列化失败",
				Data: "",
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func (jrAPI *JRAPI) ApplyClaim(c *gin.Context) {
	var reApplyClaim jrrequest.JRAPIApplyClaim
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reApplyClaim)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resApplyClaim jrresponse.JRAPIApplyClaim
	if resApplyClaim, err = jrAPIService.ApplyClaim(reApplyClaim); err != nil {
		global.GVA_LOG.Error("创建Claim失败!", zap.Error(err))
		if err.Error() == jrapi.MissingServiceParam.String() {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingServiceParam),
				Msg:  jrapi.MissingServiceParam.String(),
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.FAILED),
				Msg:  err.Error(),
				Data: "",
			})
		}
	} else {
		var response jrresponse.JRResponse
		if response, err = lg.GenJRResponse(resApplyClaim); err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: -1,
				Msg:  "resApply序列化失败",
				Data: "",
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func (jrAPI *JRAPI) LogoutPush(c *gin.Context) {
	var reLogoutPush jrrequest.JRAPILogoutPush
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reLogoutPush)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resLogoutPush jrresponse.JRAPILogoutPush
	if resLogoutPush, err = jrAPIService.LogoutPush(reLogoutPush); err != nil {
		global.GVA_LOG.Error("创建Logout失败!", zap.Error(err))
		if err.Error() == jrapi.MissingServiceParam.String() {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingServiceParam),
				Msg:  jrapi.MissingServiceParam.String(),
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.FAILED),
				Msg:  err.Error(),
				Data: "",
			})
		}
	} else {
		var response jrresponse.JRResponse
		if response, err = lg.GenJRResponse(resLogoutPush); err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: -1,
				Msg:  "resApply序列化失败",
				Data: "",
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func (jrAPI *JRAPI) ApplyInvoice(c *gin.Context) {
	var reApplyInvoice jrrequest.JRAPIApplyInvoice
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reApplyInvoice)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resApplyInvoice jrresponse.JRAPIApplyInvoice
	if resApplyInvoice, err = jrAPIService.ApplyInvoice(reApplyInvoice); err != nil {
		global.GVA_LOG.Error("创建ApplyInvoice失败!", zap.Error(err))
		if err.Error() == jrapi.MissingServiceParam.String() {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingServiceParam),
				Msg:  jrapi.MissingServiceParam.String(),
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.FAILED),
				Msg:  err.Error(),
				Data: "",
			})
		}
	} else {
		var response jrresponse.JRResponse
		if response, err = lg.GenJRResponse(resApplyInvoice); err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: -1,
				Msg:  "resApply序列化失败",
				Data: "",
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func (jrAPI *JRAPI) LetterFileDownload(c *gin.Context) {
	elog, ok := c.GetQuery("elog")
	if !ok {
		c.String(http.StatusOK, "非法参数")
		return
	}
	encrypt, ok := c.GetQuery("type")
	encryptFlag := false
	if ok {
		if encrypt == "encrypt" {
			encryptFlag = true
		} else {
			c.String(http.StatusOK, "非法参数")
			return
		}
	}

	if refile, err := jrAPIService.LetterFileDownload(elog, encryptFlag); err != nil {
		c.String(http.StatusOK, "获取文件失败")
		return
	} else {
		c.Header("Content-Disposition", "attachment; filename="+*refile.FileName) // 用来指定下载下来的文件名
		c.Data(http.StatusOK, "application/octet-stream", refile.FileSteam)
	}
}

func (jrAPI *JRAPI) DelayFileDownload(c *gin.Context) {
	elog, ok := c.GetQuery("elog")
	if !ok {
		c.String(http.StatusOK, "非法参数")
		return
	}
	encrypt, ok := c.GetQuery("type")
	encryptFlag := false
	if ok {
		if encrypt == "encrypt" {
			encryptFlag = true
		} else {
			c.String(http.StatusOK, "非法参数")
			return
		}
	}

	if refile, err := jrAPIService.DelayFileDownload(elog, encryptFlag); err != nil {
		c.String(http.StatusOK, "获取文件失败")
		return
	} else {
		c.Header("Content-Disposition", "attachment; filename="+*refile.FileName) // 用来指定下载下来的文件名
		c.Data(http.StatusOK, "application/octet-stream", refile.FileSteam)
	}
}

func (jrAPI *JRAPI) InvoiceFileDownload(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
