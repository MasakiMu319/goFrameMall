package backend

import "github.com/gogf/gf/v2/frame/g"

type PermissionCreateUpdateBase struct {
	Name string `json:"name" v:"required#权限名不能为空" dc:"权限名称"`
	Path string `json:"path"  dc:"权限路径"`
}

type PermissionReq struct {
	g.Meta `path:"/backend/permission/add" tags:"permission" method:"post" summary:"添加权限"`
	PermissionCreateUpdateBase
}

type PermissionRes struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionUpdateReq struct {
	g.Meta `path:"/backend/permission/update" tags:"permission" method:"post" summary:"修改权限"`
	Id     uint `json:"id" v:"required#权限名不能为空" dc:"权限 id"`
	PermissionCreateUpdateBase
}

type PermissionUpdateRes struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionDeleteReq struct {
	g.Meta `path:"/backend/permission/delete" method:"delete" tags:"管理员" summary:"删除权限"`
	Id     uint `v:"min:1#请选择需要删除的权限" dc:"权限 id"`
}
type PermissionDeleteRes struct{}

type PermissionGetListCommonReq struct {
	g.Meta `path:"/backend/permission/list" method:"get" tags:"权限列表" summary:"获取权限列表接口"`
	CommonPaginationReq
}
type PermissionGetListCommonRes struct {
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	// TODO: 前后端分离项目不能直接返回 HTML
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
