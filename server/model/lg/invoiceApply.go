package lg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type InvoiceApply struct {
	global.GVA_MODEL
	ApplyNo            *string  `json:"applyNo" form:"applyNo" gorm:"type:varchar(64);"`
	InvoiceTotalAmount *float64 `json:"invoiceTotalAmount" form:"invoiceTotalAmount"`
	InvoiceType        *string  `json:"invoiceType" form:"invoiceType" gorm:"type:varchar(2);"`
	InvoiceTileType    *string  `json:"invoiceTileType" form:"invoiceTileType" gorm:"type:varchar(2);"`
	InvoiceTile        *string  `json:"invoiceTile" form:"invoiceTile" gorm:"type:varchar(256);"`
	TaxNo              *string  `json:"taxNo" form:"taxNo" gorm:"type:varchar(20);"`
	BankName           *string  `json:"bankName" form:"bankName" gorm:"type:varchar(256);"`
	BankNo             *string  `json:"bankNo" form:"bankNo" gorm:"type:varchar(64);"`
	CompanyAddress     *string  `json:"companyAddress" form:"companyAddress" gorm:"type:varchar(512);"`
	CompanyTel         *string  `json:"companyTel" form:"companyTel" gorm:"type:varchar(64);"`
	Remarks            *string  `json:"remarks" form:"remarks" gorm:"type:varchar(1024);"`
	OrderList          *string  `json:"orderList" form:"orderList" gorm:"type:text;"`
	AuditStatus        *int64   `json:"auditStatus" form:"auditStatus"`
	AuditOpinion       *string  `json:"auditOpinion" form:"auditOpinion" gorm:"type:varchar(512);"`
	AuditDate          *string  `json:"auditDate" form:"auditDate" gorm:"type:varchar(19);"`
	InvoiceList        *string  `json:"invoiceList" form:"invoiceList" gorm:"type:text;"`
}

func (InvoiceApply) TableName() string {
	return "lg_invoice_apply"
}
