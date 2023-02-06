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
	g.Meta `path:"/backend/rotation/delete" method:"delete" tags:"内容" summary:"删除轮播图"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type RotationDeleteRes struct{}
