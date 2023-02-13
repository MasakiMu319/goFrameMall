package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/backend/role/add" tags:"role" method:"post" summary:"添加角色"`
	Name   string `json:"name" v:"required#角色名不能为空" dc:"角色名称"`
	Desc   string `json:"desc"  dc:"角色描述"`
}

type RoleRes struct {
	RoleId uint `json:"role_id"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update" tags:"role" method:"post" summary:"修改角色"`
	Id     uint   `json:"id" v:"required#角色名不能为空" dc:"角色 id"`
	Name   string `json:"name" v:"required#角色名不能为空" dc:"角色名称"`
	Desc   string `json:"desc"  dc:"角色描述"`
}

type RoleUpdateRes struct {
	RoleId uint `json:"role_id"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete" method:"delete" tags:"管理员" summary:"删除角色"`
	Id     uint `v:"min:1#请选择需要删除的角色" dc:"角色 id"`
}
type RoleDeleteRes struct{}

type RoleGetListCommonReq struct {
	g.Meta `path:"/backend/role/list" method:"get" tags:"角色列表" summary:"获取角色列表接口"`
	CommonPaginationReq
}
type RoleGetListCommonRes struct {
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	// TODO: 前后端分离项目不能直接返回 HTML
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type AddPermissionReq struct {
	g.Meta       `path:"/backend/role/add/permission" method:"post" tags:"角色" summary:"角色添加权限"`
	RoleId       uint `json:"role_id"  desc:"角色 id"`
	PermissionId uint `json:"permission_id" desc:"权限 id"`
}

type AddPermissionRes struct {
	Id uint `json:"id"`
}

type DeletePermissionReq struct {
	g.Meta       `path:"/backend/role/delete/permission" method:"post" tags:"角色" summary:"角色删除权限"`
	RoleId       uint `json:"role_id"  desc:"角色 id"`
	PermissionId uint `json:"permission_id" desc:"权限 id"`
}

type DeletePermissionRes struct {
}
