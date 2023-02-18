package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Invoice struct {
	global.GVA_MODEL
	InvoiceNo          *string  `json:"invoiceNo" form:"invoiceNo" gorm:"type:varchar(64);"`
	InvoiceAmount      *float64 `json:"invoiceAmount" form:"invoiceAmount"`
	InvoiceType        *string  `json:"invoiceType" form:"invoiceType" gorm:"type:varchar(2);"`
	InvoiceForm        *string  `json:"invoiceForm" form:"invoiceForm" gorm:"type:varchar(2);"`
	InvoicePoint       *float64 `json:"invoicePoint" form:"invoicePoint"`
	InvoiceContent     *string  `json:"invoiceContent" form:"invoiceContent" gorm:"type:varchar(256);"`
	InvoiceRemark      *string  `json:"invoiceRemark" form:"invoiceRemark" gorm:"type:varchar(256);"`
	InvoiceTime        *string  `json:"invoiceTime" form:"invoiceTime" gorm:"type:varchar(19);"`
	InvoiceDownloadUrl *string  `json:"invoiceDownloadUrl" form:"invoiceDownloadUrl" gorm:"type:varchar(256);"`
	OrderList          *string  `json:"orderList" form:"orderList" gorm:"type:text;"`
}

func (Invoice) TableName() string {
	return "lg_invoice"
}
