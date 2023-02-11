package role

import (
	"context"
	"goFrameMall/internal/dao"
	"goFrameMall/internal/model"
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
