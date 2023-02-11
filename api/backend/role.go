package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/backend/role/add" tags:"role" method:"post" summary:"添加角色"`
	Name   string `json:"name" v:"required#角色名不能为空" dc:"角色名称"`
	Desc   string `json:"desc"  dc:"角色描述"`
}

type RoleRes struct {
	RoleId int `json:"role_id"`
}
