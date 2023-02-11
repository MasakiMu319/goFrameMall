package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"goFrameMall/internal/dao"
	"goFrameMall/internal/model"
	"goFrameMall/internal/service"
	"goFrameMall/utility"
)

type sData struct {
}

func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {
	return &model.DataHeadOutput{
		TodayOrderCount: todayOderCount(ctx),
		// TODO: 日活暂时用随机数处理
		DAU:            utility.RandInt(1000),
		ConversionRate: utility.RandInt(50),
	}, nil
}

func todayOderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).
		Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return count
}
