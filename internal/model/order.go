package model

import "github.com/gogf/gf/v2/os/gtime"

// OrderCreateGetBase 创建/修改订单基类
type OrderCreateGetBase struct {
	//Number           string `orm:"number" json:"number" description:"订单编号"`
	Number           string
	UserId           uint
	PayType          int
	Remark           string
	PayAt            string
	Status           int
	ConsigneeName    string
	ConsigneePhone   string
	ConsigneeAddress string
	Price            int
	Coupon           int
	ActualPrice      int
}

// OrderCreateInput 创建订单
type OrderCreateInput struct {
	OrderCreateGetBase
	GoodsList []OrderGoodsGetUpdateBase
}

// OrderCreateOutput 创建订单返回结果
type OrderCreateOutput struct {
	OrderId int `json:"order_id"`
}

type OrderGoodsGetUpdateBase struct {
	OrderId     uint
	GoodsId     uint
	Count       int
	PayType     int
	Remark      string
	Status      int
	Price       int
	CouponPrice int
	ActualPrice int
}

// OrderUpdateInput 修改订单
type OrderUpdateInput struct {
	OrderGoodsGetUpdateBase
}

// OrderGetListInput 获取订单列表
type OrderGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// OrderGetListOutput 查询列表结果
type OrderGetListOutput struct {
	List  []OrderGetListOutputItem
	Page  int
	Size  int
	Total int
}

type OrderGetListOutputItem struct {
	Id uint `json:"id"` // 自增ID
	// the nested struct can't have tag `json`, refer: https://goframe.org/pages/viewpage.action?pageId=61149338
	Order     OrderCreateGetBase
	GoodsList []OrderGoodsGetUpdateBase
	// TODO: bind goods info
	CreatedAt *gtime.Time
	UpdatedAt *gtime.Time
}

type OrderGetDetailInput struct {
	Id uint
}

type OrderGetDetailOutput struct {
	Id uint
	OrderCreateGetBase
	GoodsList []OrderGoodsGetUpdateBase
}

type OrderChangeStatusInput struct {
	AdminID uint
	Id      uint
	Status  int
}

type OrderChangeStatusOutput struct {
	Id uint
}
