package nnrequest

import "github.com/flipped-aurora/gin-vue-admin/server/model/lg"

type NNQueryInvoice struct {
	Order lg.Order `json:"order"`
}
