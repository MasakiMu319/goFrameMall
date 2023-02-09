package admin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"goFrameMall/internal/dao"
	"goFrameMall/internal/model"
	"goFrameMall/internal/model/entity"
	"goFrameMall/internal/service"
	"goFrameMall/utility"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

func (s *sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	// 处理加密盐和密码的逻辑
	userSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, userSalt)
	in.UserSalt = userSalt
	lastInsertID, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AdminCreateOutput{AdminId: int(lastInsertID)}, err
}

func (s *sAdmin) Delete(ctx context.Context, id uint) error {
	// TODO: func 部分的 tx 传递的非指针，暂不明有何影响
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.AdminInfo.Ctx(ctx).Where(g.Map{
			dao.AdminInfo.Columns().Id: id,
			//}).Delete()
			// 软删除
		}).Unscoped().Delete()
		// 使用 Unscoped() 会忽略掉数据表中 deleted_at 字段，此时是真正的物理删除
		return err
	})
}

// Update 修改
func (s *sAdmin) Update(ctx context.Context, in model.AdminUpdateInput) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		if in.Password != "" {
			userSalt := grand.S(10)
			in.Password = utility.EncryptPassword(in.Password, userSalt)
			in.UserSalt = userSalt
		}
		_, err := dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			// FieldsEX 这里是重点
			FieldsEx(dao.AdminInfo.Columns().Id).
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error) {
	var (
		m = dao.AdminInfo.Ctx(ctx)
	)
	out = &model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.AdminInfo
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
	// Admin
	// TODO: ScanList 是从 List 这个 AdminGetListOutputItem 结构中取出 Admin 字段
	//if err := listModel.ScanList(&out.List, "Admin"); err != nil {
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (s *sAdmin) GetAdminByNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{} {
	// TODO: 对接 DB
	//if in.Name == "admin" && in.Password == "admin" {
	//	return g.Map{
	//		"id":       1,
	//		"username": "admin",
	//	}
	//}
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return nil
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	} else {
		return g.Map{
			"id":       adminInfo.Id,
			"username": adminInfo.Name,
		}
	}
}
