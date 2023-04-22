package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	lgReq "github.com/flipped-aurora/gin-vue-admin/server/model/lg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type OrderApi struct {
}

var orderService = service.ServiceGroupApp.LgServiceGroup.OrderService

func (orderApi *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo lgReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderService.GetOrderInfoList(pageInfo); err != nil {
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

func (orderApi *OrderApi) ApproveApply(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.ApproveApply(order); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败："+err.Error(), c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (orderApi *OrderApi) RejectApply(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.RejectApply(order); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败："+err.Error(), c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (orderApi *OrderApi) ApproveDelay(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.ApproveDelay(order); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (orderApi *OrderApi) RejectDelay(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.RejectDelay(order); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (orderApi *OrderApi) ApproveClaim(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.ApproveClaim(order); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (orderApi *OrderApi) RejectClaim(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.RejectClaim(order); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (orderApi *OrderApi) OpenLetter(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.OpenLetter(order); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (orderApi *OrderApi) RePush(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.RePush(order); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (orderApi *OrderApi) GetInsuranceBalance(c *gin.Context) {
	if insuranceBalance, err := orderService.GetInsuranceBalance(); err != nil {
		global.GVA_LOG.Error("获取额度失败!", zap.Error(err))
		response.FailWithMessage("获取额度失败："+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{
			"insuranceBalance": insuranceBalance,
		}, "获取成功", c)
	}
}

func (orderApi *OrderApi) GetOrderStatisticData(c *gin.Context) {
	var pageInfo lgReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if orderStatisticData, err := orderService.GetOrderStatisticData(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{
			"orderStatisticData": orderStatisticData,
		}, "获取成功", c)
	}
}

func (orderApi *OrderApi) GetEmployeeStatisticData(c *gin.Context) {
	if employeeStatisticData, err := orderService.GetEmployeeStatisticData(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{
			"employeeStatisticData": employeeStatisticData,
		}, "获取成功", c)
	}
}

func (orderApi *OrderApi) GetGEODistributionData(c *gin.Context) {
	var pageInfo lgReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if geoDistributionData, err := orderService.GetGEODistributionData(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{
			"geoDistributionData": geoDistributionData,
		}, "获取成功", c)
	}
}

func (orderApi *OrderApi) GetOrderTrendData(c *gin.Context) {
	var pageInfo lgReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if orderTrendData, err := orderService.GetOrderTrendData(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{
			"orderTrendData": orderTrendData,
		}, "获取成功", c)
	}
}

func (orderApi *OrderApi) ExportExcel(c *gin.Context) {
	var pageInfo lgReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if excel, err := orderService.ExportExcel(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		c.Writer.Header().Add("success", "true")
		c.Header("Content-Disposition", "attachment; filename="+strconv.Itoa(int(time.Now().Unix()))+".xlsx") // 用来指定下载下来的文件名
		c.Data(http.StatusOK, "application/octet-stream", excel)
	}
}

func (orderApi *OrderApi) AssignOrder(c *gin.Context) {
	var assign lgReq.AssignOrder
	err := c.ShouldBindJSON(&assign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.AssignOrder(assign); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败："+err.Error(), c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (orderApi *OrderApi) MarkOfflineRefund(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.MarkOfflineRefund(order); err != nil {
		global.GVA_LOG.Error("标记失败!", zap.Error(err))
		response.FailWithMessage("标记失败："+err.Error(), c)
	} else {
		response.OkWithMessage("绑定成功", c)
	}
}

func (orderApi *OrderApi) UnmarkOfflineRefund(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.UnmarkOfflineRefund(order); err != nil {
		global.GVA_LOG.Error("标记失败!", zap.Error(err))
		response.FailWithMessage("标记失败："+err.Error(), c)
	} else {
		response.OkWithMessage("标记成功", c)
	}
}

func (orderApi *OrderApi) ExportInvoiceExcel(c *gin.Context) {
	var pageInfo lgReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if excel, err := orderService.ExportInvoiceExcel(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		c.Writer.Header().Add("success", "true")
		c.Header("Content-Disposition", "attachment; filename="+strconv.Itoa(int(time.Now().Unix()))+".xlsx") // 用来指定下载下来的文件名
		c.Data(http.StatusOK, "application/octet-stream", excel)
	}
}

func (orderApi *OrderApi) ElogValidate(c *gin.Context) {
	var elogValidate lgReq.ElogValidate
	err := c.ShouldBindJSON(&elogValidate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if elogValidateMessage, err := orderService.ElogValidate(elogValidate); err != nil {
		global.GVA_LOG.Error("鉴真失败!", zap.Error(err))
		response.FailWithMessage("鉴真失败："+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{
			"elogValidateMessage": elogValidateMessage,
		}, "鉴真成功", c)
	}
}
