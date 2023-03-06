package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type OrderBase struct {
	Number           string `json:"number"   dc:"订单编号"`
	UserId           uint   `json:"user_id"    v:"required#用户id不能为空" dc:"用户id"`
	PayType          int    `json:"pay_type"    v:"required#支付类型不能为空" dc:"支付类型"`
	Remark           string `json:"remark"    dc:"备注"`
	PayAt            string `json:"pay_at"    dc:"支付时间"` // pay_at allow empty, the status is must be 0
	Status           int    `json:"status"    v:"required#订单状态不能为空" dc:"状态"`
	ConsigneeName    string `json:"consignee_name"    v:"required#收货人姓名不能为空" dc:"收货人姓名"`
	ConsigneePhone   string `json:"consignee_phone"    v:"required#收货人电话不能为空" dc:"收货人电话"`
	ConsigneeAddress string `json:"consignee_address"    v:"required#收货人地址不能为空" dc:"收货人地址"`
	Price            int    `json:"price"    v:"required#订单价格不能为空" dc:"订单价格"`
	Coupon           int    `json:"coupon"   dc:"优惠券价格"`
	ActualPrice      int    `json:"actual_price"    v:"required#实际支付价格不能为空" dc:"实际支付价格"`
}

// OrderReq Order 请求
type OrderReq struct {
	g.Meta `path:"/backend/order/add" tags:"Order" method:"post"`
	OrderBase
	GoodsList []OrderGoodsBase `json:"goods_list"    v:"required#商品列表不能为空" dc:"商品列表"`
}

// OrderRes Order 响应
type OrderRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	OrderId int `json:"orderId"`
}

type OrderGoodsBase struct {
	GoodsId     uint   `json:"goods_id"    v:"required#商品id不能为空" dc:"商品id"`
	Count       int    `json:"count"    v:"required#商品数量不能为空" dc:"商品数量"`
	PayType     int    `json:"pay_type"    v:"required#支付类型不能为空" dc:"支付类型"`
	Remark      string `json:"remark"    dc:"备注"`
	Price       int    `json:"price"    v:"required#订单价格不能为空" dc:"订单价格"`
	CouponPrice int    `json:"coupon_price"    dc:"优惠券价格"`
	ActualPrice int    `json:"actual_price"    v:"required#实际支付价格不能为空" dc:"实际支付价格"`
}

type OrderUpdateReq struct {
	g.Meta `path:"/backend/order/update/{id}" method:"post" tags:"订单" summary:"修改订单接口"`
	// TODO: is there need to add a validator for userId?
	Id uint `json:"id"      v:"min:1#请选择需要修改的订单" dc:"订单Id"`
	OrderGoodsBase
}

type OrderUpdateRes struct {
	OrderId uint `json:"order_id"`
}

type OrderGetListCommonReq struct {
	g.Meta `path:"/backend/order/list" method:"get" tags:"订单列表" summary:"获取订单列表接口"`
	CommonPaginationReq
}

type OrderGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// TODO: add a modify order status api

type OrderGetDetailReq struct {
	g.Meta `path:"/backend/order/detail/{id}" method:"get" tags:"订单详情" summary:"获取订单详情接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要查看的订单" dc:"订单Id"`
}

type OrderGetDetailRes struct {
	Id uint `json:"id"`
	OrderBase
	GoodsList []OrderGoodsBase `json:"goods_list"`
}

type OrderChangeStatusReq struct {
	g.Meta  `path:"/backend/order/status/{id}" method:"post" tags:"订单状态" summary:"修改订单状态接口"`
	AdminId uint `json:"admin_id"      v:"min:1#管理员 id" dc:"管理员 id"`
	Id      uint `json:"id"      v:"min:1#请选择需要修改的订单" dc:"订单Id"`
	Status  int  `json:"status"    v:"required#订单状态不能为空" dc:"状态"`
}

type OrderChangeStatusRes struct {
	OrderId uint `json:"order_id"`
}
