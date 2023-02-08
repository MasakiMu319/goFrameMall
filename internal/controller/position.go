package controller

import (
	"context"
	"goFrameMall/api/backend"
	"goFrameMall/internal/model"
	"goFrameMall/internal/service"
)

// Position 内容管理
var Position = cPosition{}

type cPosition struct{}

func (a *cPosition) Create(ctx context.Context, req *backend.PositionReq) (res *backend.PositionRes, err error) {
	out, err := service.Position().Create(ctx, model.PositionCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
			Sort:      req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionRes{PositionId: out.PositionId}, nil
}

func (a *cPosition) Delete(ctx context.Context, req *backend.PositionDeleteReq) (res *backend.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.Id)
	return
}

func (a *cPosition) Update(ctx context.Context, req *backend.PositionUpdateReq) (res *backend.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, model.PositionUpdateInput{
		Id: req.Id,
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
			Link:      req.Link,
			Sort:      req.Sort,
		},
	})
	return
}

// List position list
func (a *cPosition) List(ctx context.Context, req *backend.PositionGetListCommonReq) (res *backend.PositionGetListCommonRes, err error) {
	getListRes, err := service.Position().GetList(ctx, model.PositionGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Total: getListRes.Total,
		Size:  getListRes.Size}, nil
}
