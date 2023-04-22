package lg

type RouterGroup struct {
	FileRouter
	InvoiceRouter
	InvoiceApplyRouter
	LetterRouter
	LogoutRouter
	OrderRouter
	PayRouter
	ProjectRouter
	RefundRouter
	RevokeRouter
	TemplateRouter
}
