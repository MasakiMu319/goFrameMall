package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RotationReq Rotation 请求
type RotationReq struct {
	// g.Meta 就是这个服务所指向的路由，method 表示这个服务使用哪种 HTTP 方法访问
	g.Meta `path:"/backend/rotation/add" tags:"Rotation" method:"post" summary:"You first Rotation api"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	// Sort 可以为空，所以不用设置 v，猜测 v 与 verify 有关，是验证为空时报错的信息
	Sort int `json:"sort"     dc:"轮播图的序号"`
}

// RotationRes Rotation 响应
type RotationRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	RotationId int `json:"rotationId"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/backend/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图"`
	Id     uint `v:"min:1#请选择需要删除的轮播图" dc:"轮播图 id"`
}
type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/backend/rotation/update/{id}" method:"post" tags:"轮播图" summary:"修改轮播图接口"`
	Id     uint   `json:"id"      v:"min:1#请选择需要修改的轮播图" dc:"轮播图Id"`
	PicUrl string `json:"pic_url"    v:"required#轮播图链接不能为空" dc:"轮播图链接"`
	Link   string `json:"link"    v:"required#轮播图跳转链接不能为空" dc:"轮播图跳转链接"`
	Sort   int    `json:"sort"    dc:"轮播图排序值"`
}
type RotationUpdateRes struct{}

type RotationGetListCommonReq struct {
	g.Meta `path:"/backend/rotation/list" method:"get" tags:"轮播图列表" summary:"获取轮播图列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type RotationGetListCommonRes struct {
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	// TODO: 前后端分离项目不能直接返回 HTML
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
