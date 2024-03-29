package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"time"
)

type ProjectSearch struct {
	lg.Project
	StartCreatedAt *time.Time  `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time  `json:"endCreatedAt" form:"endCreatedAt"`
	OpenTime       []time.Time `json:"openTime" form:"openTime[]"`
	request.PageInfo
}
