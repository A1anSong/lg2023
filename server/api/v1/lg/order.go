package lg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/nn/nnrequest"
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

func (orderApi *OrderApi) CreateOrder(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.CreateOrder(order); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (orderApi *OrderApi) DeleteOrder(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.DeleteOrder(order); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (orderApi *OrderApi) DeleteOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.DeleteOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (orderApi *OrderApi) UpdateOrder(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.UpdateOrder(order); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (orderApi *OrderApi) FindOrder(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindQuery(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorder, err := orderService.GetOrder(order.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorder": reorder}, c)
	}
}

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

func (orderApi *OrderApi) ApproveRefund(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.ApproveRefund(order); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

func (orderApi *OrderApi) RejectRefund(c *gin.Context) {
	var order lg.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.RejectRefund(order); err != nil {
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
	var pageInfo lgReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if employeeStatisticData, err := orderService.GetEmployeeStatisticData(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{
			"employeeStatisticData": employeeStatisticData,
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

func (orderApi *OrderApi) FindOrderByNos(c *gin.Context) {
	type OrderByNos struct {
		OrderNos []string `json:"orderNos[]" form:"orderNos[]"`
	}
	var orderByNos OrderByNos
	err := c.ShouldBindQuery(&orderByNos)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if orders, err := orderService.GetOrderByNos(orderByNos.OrderNos); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"orders": orders}, c)
	}
}

func (orderApi *OrderApi) RequestInvoice(c *gin.Context) {
	var reqInvoice nnrequest.NNRequestInvoice
	err := c.ShouldBindJSON(&reqInvoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.RequestInvoice(reqInvoice); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功，等待约一分钟左右可点击查询开票结果", c)
	}
}

func (orderApi *OrderApi) QueryInvoice(c *gin.Context) {
	var reqInvoice nnrequest.NNQueryInvoice
	err := c.ShouldBindJSON(&reqInvoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.QueryInvoice(reqInvoice); err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
	} else {
		response.OkWithMessage("提交成功，等待约一分钟左右可点击查询开票结果", c)
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
