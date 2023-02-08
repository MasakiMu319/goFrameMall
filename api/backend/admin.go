package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdminReq Admin 请求
type AdminReq struct {
	// g.Meta 就是这个服务所指向的路由，method 表示这个服务使用哪种 HTTP 方法访问
	g.Meta   `path:"/backend/admin/add" tags:"Admin" method:"post" summary:"You first Admin api"`
	Name     string `json:"name"    v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids"    dc:"角色 ids"`
	IsAdmin  int    `json:"is_admin"    dc:"是否超级管理员"`
}

// AdminRes Admin 响应
type AdminRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	AdminId int `json:"adminId"`
}

type AdminDeleteReq struct {
	g.Meta `path:"/backend/admin/delete" method:"delete" tags:"管理员" summary:"删除管理员"`
	Id     uint `v:"min:1#请选择需要删除的管理员" dc:"管理员 id"`
}
type AdminDeleteRes struct{}

type AdminUpdateReq struct {
	g.Meta   `path:"/backend/admin/update/{id}" method:"post" tags:"管理员" summary:"修改管理员接口"`
	Id       uint   `json:"id"      v:"min:1#请选择需要修改的管理员" dc:"管理员Id"`
	Name     string `json:"name"    v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids"    dc:"角色 ids"`
	IsAdmin  int    `json:"is_admin"    dc:"是否超级管理员"`
}
type AdminUpdateRes struct {
	AdminId uint `json:"admin_id"`
}

type AdminGetListCommonReq struct {
	g.Meta `path:"/backend/admin/list" method:"get" tags:"管理员列表" summary:"获取管理员列表接口"`
	CommonPaginationReq
}
type AdminGetListCommonRes struct {
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	// TODO: 前后端分离项目不能直接返回 HTML
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
