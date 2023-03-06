package order

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"goFrameMall/internal/dao"
	"goFrameMall/internal/model"
	"goFrameMall/internal/model/entity"
	"goFrameMall/internal/service"
	"goFrameMall/utility"
	"log"
	"time"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

func (s *sOrder) Create(ctx context.Context, in model.OrderCreateInput) (out *model.OrderCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	// add order number
	in.OrderCreateGetBase.Number = utility.GetOderNumber()
	// TODO:
	if in.PayAt == "" {
		in.PayAt = utility.TimeStampToDateTime(time.Now().Unix())
	}
	// TODO: we do not update order goods info here
	lastInsertID, err := dao.OrderInfo.Ctx(ctx).Data(in.OrderCreateGetBase).InsertAndGetId()
	log.Println(in.GoodsList)
	for _, goods := range in.GoodsList {
		goods.OrderId = uint(lastInsertID)
		goods.PayType = in.PayType
		//goods.Status = in.Status
		_, err := dao.OrderGoodsInfo.Ctx(ctx).Data(goods).Insert()
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return out, err
	}
	return &model.OrderCreateOutput{OrderId: int(lastInsertID)}, err
}

// Update 修改
func (s *sOrder) Update(ctx context.Context, in model.OrderUpdateInput) error {
	return dao.OrderGoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		// update order info
		// only permit to update order status is 0 (unpaid)
		// TODO: the question is this will update pay_at
		_, err := dao.OrderGoodsInfo.
			Ctx(ctx).
			Data(in).
			// TODO: there may be a bug
			FieldsEx(dao.OrderGoodsInfo.Columns().OrderId,
				dao.OrderGoodsInfo.Columns().Id,
				dao.OrderGoodsInfo.Columns().GoodsId,
				dao.OrderGoodsInfo.Columns().Status,
				dao.OrderGoodsInfo.Columns().PayAt).
			Where(dao.OrderGoodsInfo.Columns().GoodsId, in.GoodsId).
			Where(dao.OrderGoodsInfo.Columns().Id, in.OrderId).
			// check order status
			Where(dao.OrderGoodsInfo.Columns().Status, 0).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sOrder) GetList(ctx context.Context, in model.OrderGetListInput) (out *model.OrderGetListOutput, err error) {
	var (
		m = dao.OrderInfo.Ctx(ctx)
	)
	out = &model.OrderGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.OrderInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Order
	// TODO: ScanList 是从 List 这个 OrderGetListOutputItem 结构中取出 Order 字段
	//if err := listModel.ScanList(&out.List, "OrderCreateGetBase"); err != nil {
	if err := listModel.Scan(&out.List); err != nil {
		return nil, err
	}
	log.Println("length: ", len(out.List))
	for i, v := range out.List {
		var order model.OrderCreateGetBase
		if err := listModel.Scan(&order); err != nil {
			return nil, err
		}
		out.List[i].Order = order
		cur := v.Id
		var goodsList []model.OrderGoodsGetUpdateBase
		if err := dao.OrderGoodsInfo.Ctx(ctx).Where(dao.OrderGoodsInfo.Columns().OrderId, cur).Scan(&goodsList); err != nil {
			return nil, err
		}
		out.List[i].GoodsList = goodsList
	}
	return
}

func (s *sOrder) GetDetail(ctx context.Context, in model.OrderGetDetailInput) (out *model.OrderGetDetailOutput, err error) {
	var (
		m = dao.OrderInfo.Ctx(ctx)
	)
	out = &model.OrderGetDetailOutput{}
	out.Id = in.Id
	// 查询
	if err := m.Where(dao.OrderInfo.Columns().Id, in.Id).Scan(&out.OrderCreateGetBase); err != nil {
		return out, err
	}
	// 查询
	if err := dao.OrderGoodsInfo.Ctx(ctx).Where(dao.OrderGoodsInfo.Columns().OrderId, in.Id).Scan(&out.GoodsList); err != nil {
		return out, err
	}
	return
}

func (s *sOrder) ChangeStatus(ctx context.Context, in model.OrderChangeStatusInput) (out *model.OrderChangeStatusOutput, err error) {
	value, err := dao.AdminInfo.Ctx(ctx).Data(in).
		Fields(dao.AdminInfo.Columns().Id, dao.AdminInfo.Columns().IsAdmin).
		Where(dao.AdminInfo.Columns().Id, in.AdminID).
		Value(dao.AdminInfo.Columns().IsAdmin)
	if err != nil {
		return nil, err
	}
	// check if admin
	if value.Int() == 0 {
		return nil, errors.New("only admin can change order status")
	}
	// update order info status
	err = dao.OrderInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.OrderInfo.Ctx(ctx).Data(in).
			Fields(dao.OrderInfo.Columns().Status).
			Where(dao.OrderInfo.Columns().Id, in.Id).
			// check order status
			Update()
		return err
	})
	if err != nil {
		return nil, err
	}
	// update order goods info status
	err = dao.OrderGoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.OrderGoodsInfo.Ctx(ctx).Data(in).
			Fields(dao.OrderGoodsInfo.Columns().Status).
			Where(dao.OrderGoodsInfo.Columns().OrderId, in.Id).
			// check order status
			Update()
		if err != nil {
			return err
		}
		_, err = dao.OrderGoodsInfo.Ctx(ctx).Data(in).
			Fields(dao.OrderGoodsInfo.Columns().Status).
			Where(dao.OrderGoodsInfo.Columns().Id, in.Id).Update()
		return err
	})
	if err != nil {
		return nil, err
	}
	return &model.OrderChangeStatusOutput{
		Id: in.Id,
	}, err
}
