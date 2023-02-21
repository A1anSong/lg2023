package nnresponse

type NNRequestInvoice struct {
	Code     string      `json:"code"`
	Describe string      `json:"describe"`
	Result   interface{} `json:"result"`
}
