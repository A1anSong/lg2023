package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Invoice struct {
	global.GVA_MODEL
	InvoiceNo          *string  `json:"invoiceNo" form:"invoiceNo" gorm:"type:varchar(64);"`
	InvoiceAmount      *float64 `json:"invoiceAmount" form:"invoiceAmount"`
	InvoiceType        *string  `json:"invoiceType" form:"invoiceType" gorm:"type:varchar(2);"`
	InvoiceTileType    *string  `json:"invoiceTileType" form:"invoiceTileType" gorm:"type:varchar(2);"`
	InvoiceTile        *string  `json:"InvoiceTile" form:"InvoiceTile" gorm:"type:varchar(200);"`
	TaxNo              *string  `json:"TaxNo" form:"TaxNo" gorm:"type:varchar(20);"`
	BankName           *string  `json:"BankName" form:"BankName" gorm:"type:varchar(200);"`
	BankNo             *string  `json:"BankNo" form:"BankNo" gorm:"type:varchar(19);"`
	CompanyAddress     *string  `json:"CompanyAddress" form:"CompanyAddress" gorm:"type:varchar(512);"`
	CompanyTel         *string  `json:"CompanyTel" form:"CompanyTel" gorm:"type:varchar(64);"`
	Remarks            *string  `json:"Remarks" form:"InvoiceTile" gorm:"type:varchar(200);"`
	InvoiceForm        *string  `json:"invoiceForm" form:"invoiceForm" gorm:"type:varchar(2);"`
	InvoicePoint       *float64 `json:"invoicePoint" form:"invoicePoint"`
	InvoiceContent     *string  `json:"invoiceContent" form:"invoiceContent" gorm:"type:varchar(256);"`
	InvoiceRemark      *string  `json:"invoiceRemark" form:"invoiceRemark" gorm:"type:varchar(256);"`
	InvoiceTime        *string  `json:"invoiceTime" form:"invoiceTime" gorm:"type:varchar(19);"`
	InvoiceSerialNum   *string  `json:"invoiceSerialNum" form:"invoiceSerialNum" gorm:"type:varchar(64);"`
	InvoiceDownloadUrl *string  `json:"invoiceDownloadUrl" form:"invoiceDownloadUrl" gorm:"type:varchar(256);"`
	OrderList          *string  `json:"orderList" form:"orderList" gorm:"type:text;"`
}

func (Invoice) TableName() string {
	return "lg_invoice"
}
