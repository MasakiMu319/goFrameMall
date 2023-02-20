package permission

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMall/internal/dao"
	"goFrameMall/internal/model"
	"goFrameMall/internal/model/entity"
	"goFrameMall/internal/service"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

func (s *sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {
	permissionId, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PermissionCreateOutput{PermissionId: uint(permissionId)}, err
}

func (s *sPermission) Delete(ctx context.Context, id uint) error {
	// TODO: func 部分的 tx 传递的非指针，暂不明有何影响
	return dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
			dao.PermissionInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// Update 修改
func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	_, err := dao.PermissionInfo.
		Ctx(ctx).
		Data(in).
		// FieldsEX 这里是重点
		FieldsEx(dao.PermissionInfo.Columns().Id).
		Where(dao.PermissionInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询内容列表
func (s *sPermission) GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error) {
	var (
		m = dao.PermissionInfo.Ctx(ctx)
	)
	out = &model.PermissionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.PermissionInfo
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
	// Permission
	// TODO: ScanList 是从 List 这个 PermissionGetListOutputItem 结构中取出 Permission 字段
	//if err := listModel.ScanList(&out.List, "Permission"); err != nil {
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
