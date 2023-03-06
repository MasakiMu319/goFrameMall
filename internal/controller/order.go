package controller

import (
	"context"
	"goFrameMall/api/backend"
	"goFrameMall/internal/model"
	"goFrameMall/internal/service"
)

// Order 内容管理
var Order = cOrder{}

type cOrder struct{}

func (a *cOrder) Create(ctx context.Context, req *backend.OrderReq) (res *backend.OrderRes, err error) {
	tmp := make([]model.OrderGoodsGetUpdateBase, 0)
	for _, v := range req.GoodsList {
		tmp = append(tmp, model.OrderGoodsGetUpdateBase{
			GoodsId:     v.GoodsId,
			Count:       v.Count,
			Remark:      v.Remark,
			Price:       v.Price,
			CouponPrice: v.CouponPrice,
			ActualPrice: v.ActualPrice,
		})
	}
	out, err := service.Order().Create(ctx, model.OrderCreateInput{
		OrderCreateGetBase: model.OrderCreateGetBase{
			UserId:           req.UserId,
			PayType:          req.PayType,
			Remark:           req.Remark,
			PayAt:            req.PayAt,
			Status:           req.Status,
			ConsigneeName:    req.ConsigneeName,
			ConsigneePhone:   req.ConsigneePhone,
			ConsigneeAddress: req.ConsigneeAddress,
			Price:            req.Price,
			Coupon:           req.Coupon,
			ActualPrice:      req.ActualPrice,
		},
		GoodsList: tmp,
	})
	if err != nil {
		return nil, err
	}
	return &backend.OrderRes{OrderId: out.OrderId}, nil
}

func (a *cOrder) Update(ctx context.Context, req *backend.OrderUpdateReq) (res *backend.OrderUpdateRes, err error) {
	err = service.Order().Update(ctx, model.OrderUpdateInput{
		OrderGoodsGetUpdateBase: model.OrderGoodsGetUpdateBase{
			OrderId:     req.Id,
			GoodsId:     req.GoodsId,
			Count:       req.Count,
			PayType:     req.PayType,
			Remark:      req.Remark,
			Price:       req.Price,
			CouponPrice: req.CouponPrice,
			ActualPrice: req.ActualPrice,
		},
	})
	return &backend.OrderUpdateRes{OrderId: req.Id}, err
}

// List order list
func (a *cOrder) List(ctx context.Context, req *backend.OrderGetListCommonReq) (res *backend.OrderGetListCommonRes, err error) {
	getListRes, err := service.Order().GetList(ctx, model.OrderGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.OrderGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Total: getListRes.Total,
		Size:  getListRes.Size}, nil
}

func (a *cOrder) GetDetail(ctx context.Context, req *backend.OrderGetDetailReq) (res *backend.OrderGetDetailRes, err error) {
	getDetailRes, err := service.Order().GetDetail(ctx, model.OrderGetDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	goods := make([]backend.OrderGoodsBase, 0)
	for _, v := range getDetailRes.GoodsList {
		goods = append(goods, backend.OrderGoodsBase{
			GoodsId:     v.GoodsId,
			Count:       v.Count,
			PayType:     v.PayType,
			Remark:      v.Remark,
			Price:       v.Price,
			CouponPrice: v.CouponPrice,
			ActualPrice: v.ActualPrice,
		})
	}
	return &backend.OrderGetDetailRes{
		Id: getDetailRes.Id,
		OrderBase: backend.OrderBase{
			Number:           getDetailRes.Number,
			UserId:           getDetailRes.UserId,
			PayType:          getDetailRes.PayType,
			Remark:           getDetailRes.Remark,
			PayAt:            getDetailRes.PayAt,
			Status:           getDetailRes.Status,
			ConsigneeName:    getDetailRes.ConsigneeName,
			ConsigneePhone:   getDetailRes.ConsigneePhone,
			ConsigneeAddress: getDetailRes.ConsigneeAddress,
			Price:            getDetailRes.Price,
			Coupon:           getDetailRes.Coupon,
			ActualPrice:      getDetailRes.ActualPrice,
		},
		GoodsList: goods,
	}, nil
}

func (a *cOrder) ChangeStatus(ctx context.Context, req *backend.OrderChangeStatusReq) (res *backend.OrderChangeStatusRes, err error) {
	changeStatusRes, err := service.Order().ChangeStatus(ctx, model.OrderChangeStatusInput{
		AdminID: req.AdminId,
		Id:      req.Id,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &backend.OrderChangeStatusRes{OrderId: changeStatusRes.Id}, nil
}
