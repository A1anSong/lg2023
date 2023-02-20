package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []system.SysBaseMenu `json:"menus"`
	AuthorityId uint                 `json:"authorityId"` // 角色ID
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		GVA_MODEL: global.GVA_MODEL{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: system.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}

func DepartmentMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		GVA_MODEL: global.GVA_MODEL{ID: 30},
		ParentId:  "0",
		Path:      "dash",
		Name:      "dash",
		Component: "view/lg/dashboard/dash.vue",
		Sort:      101,
		Meta: system.Meta{
			Title: "总览",
			Icon:  "odometer",
		},
	}, {
		GVA_MODEL: global.GVA_MODEL{ID: 11},
		ParentId:  "0",
		Path:      "person",
		Name:      "person",
		Component: "view/person/person.vue",
		Sort:      4,
		Meta: system.Meta{
			Title: "个人信息",
			Icon:  "message",
		},
	}}
}
