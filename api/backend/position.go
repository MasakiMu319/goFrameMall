package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PositionReq Position 请求
type PositionReq struct {
	// g.Meta 就是这个服务所指向的路由，method 表示这个服务使用哪种 HTTP 方法访问
	g.Meta    `path:"/backend/position/add" tags:"Position" method:"post" summary:"You first Position api"`
	PicUrl    string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   int    `json:"goods_id"  v:"required#商品 id 不能为空"   dc:"商品 id"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	// Sort 可以为空，所以不用设置 v，猜测 v 与 verify 有关，是验证为空时报错的信息
	Sort int `json:"sort"     dc:"手工位的序号"`
}

// PositionRes Position 响应
type PositionRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	PositionId int `json:"positionId"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/backend/position/delete" method:"delete" tags:"手工位" summary:"删除手工位"`
	Id     uint `v:"min:1#请选择需要删除的手工位" dc:"手工位 id"`
}
type PositionDeleteRes struct{}

type PositionUpdateReq struct {
	g.Meta    `path:"/backend/position/update/{id}" method:"post" tags:"手工位" summary:"修改手工位接口"`
	Id        uint   `json:"id"      v:"min:1#请选择需要修改的手工位" dc:"手工位Id"`
	PicUrl    string `json:"pic_url"    v:"required#手工位链接不能为空" dc:"手工位链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   int    `json:"goods_id"  v:"required#商品 id 不能为空"   dc:"商品 id"`
	Link      string `json:"link"    v:"required#手工位跳转链接不能为空" dc:"手工位跳转链接"`
	Sort      int    `json:"sort"    dc:"手工位排序值"`
}
type PositionUpdateRes struct{}

type PositionGetListCommonReq struct {
	g.Meta `path:"/backend/position/list" method:"get" tags:"手工位列表" summary:"获取手工位列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type PositionGetListCommonRes struct {
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	// TODO: 前后端分离项目不能直接返回 HTML
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
