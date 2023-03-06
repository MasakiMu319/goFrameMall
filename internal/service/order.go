// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goFrameMall/internal/model"
)

type (
	IOrder interface {
		Create(ctx context.Context, in model.OrderCreateInput) (out *model.OrderCreateOutput, err error)
		Update(ctx context.Context, in model.OrderUpdateInput) error
		GetList(ctx context.Context, in model.OrderGetListInput) (out *model.OrderGetListOutput, err error)
		GetDetail(ctx context.Context, in model.OrderGetDetailInput) (out *model.OrderGetDetailOutput, err error)
		ChangeStatus(ctx context.Context, in model.OrderChangeStatusInput) (out *model.OrderChangeStatusOutput, err error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
