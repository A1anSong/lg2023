package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/service/lg"

	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.GVA_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.GVA_CONFIG.Timer.Detail[i])
		}
	}
	if global.GVA_CONFIG.Mysql.Path == "127.0.0.1" {
		_, err := global.GVA_Timer.AddTaskByFunc("AutoMaticQueryInvoiceResult", "* * * * *", lg.AutoMaticQueryInvoiceResult)
		if err != nil {
			fmt.Println("设置每1分钟查询开票结果失败：", err)
		}
		_, err = global.GVA_Timer.AddTaskByFunc("AutoMaticAuditInvoiceApply", "*/15 * * * *", lg.AutoMaticAuditInvoiceApply)
		if err != nil {
			fmt.Println("设置每十五分钟发票申请审查失败：", err)
		}
		_, err = global.GVA_Timer.AddTaskByFunc("AutoMaticUnEnableProject", "0 * * * *", lg.AutoMaticUnEnableProject)
		if err != nil {
			fmt.Println("设置每小时自动下架项目失败：", err)
		}
		_, err = global.GVA_Timer.AddTaskByFunc("AutoMaticRefundOrder", "* * * * *", lg.AutoMaticRefundOrder)
		if err != nil {
			fmt.Println("设置每分钟自动退函失败：", err)
		}
	}
}
