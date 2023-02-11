package model

// RoleCreateUpdateBase 创建/修改角色基类
type RoleCreateUpdateBase struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// RoleCreateInput 创建角色
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// RoleCreateOutput 创建角色返回结果
type RoleCreateOutput struct {
	RoleId int `json:"role_id"`
}
