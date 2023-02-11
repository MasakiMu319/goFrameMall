package role

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMall/internal/dao"
	"goFrameMall/internal/model"
	"goFrameMall/internal/model/entity"
	"goFrameMall/internal/service"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	roleId, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{RoleId: int(roleId)}, err
}

func (s *sRole) Delete(ctx context.Context, id uint) error {
	// TODO: func 部分的 tx 传递的非指针，暂不明有何影响
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.RoleInfo.Ctx(ctx).Where(g.Map{
			dao.RoleInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// Update 修改
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	_, err := dao.RoleInfo.
		Ctx(ctx).
		Data(in).
		// FieldsEX 这里是重点
		FieldsEx(dao.RoleInfo.Columns().Id).
		Where(dao.RoleInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询内容列表
func (s *sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	var (
		m = dao.RoleInfo.Ctx(ctx)
	)
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.RoleInfo
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
	// Role
	// TODO: ScanList 是从 List 这个 RoleGetListOutputItem 结构中取出 Role 字段
	//if err := listModel.ScanList(&out.List, "Role"); err != nil {
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
