package controller

import (
	"context"
	"goFrameMall/api/backend"
	"goFrameMall/internal/model"
	"goFrameMall/internal/service"
)

// Role 管理
var Role = cRole{}

type cRole struct{}

func (c *cRole) Create(ctx context.Context, req *backend.RoleReq) (res *backend.RoleRes, err error) {
	out, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleRes{RoleId: out.RoleId}, nil
}

func (c *cRole) AddPermission(ctx context.Context, req *backend.AddPermissionReq) (res *backend.AddPermissionRes, err error) {
	permission, err := service.Role().AddPermission(ctx, model.RoleAddPermissionInput{
		PermissionId: req.PermissionId,
		RoleId:       req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.AddPermissionRes{
		Id: permission.Id,
	}, nil
}

func (c *cRole) Delete(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	return
}

func (c *cRole) DeletePermission(ctx context.Context, req *backend.DeletePermissionReq) (res *backend.DeletePermissionRes, err error) {
	err = service.Role().DeletePermission(ctx, model.RoleDeletePermissionInput{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.DeletePermissionRes{}, nil
}

func (c *cRole) Update(ctx context.Context, req *backend.RoleUpdateReq) (res *backend.RoleUpdateRes, err error) {
	err = service.Role().Update(ctx, model.RoleUpdateInput{
		Id: req.Id,
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	return
}

// List role list
func (c *cRole) List(ctx context.Context, req *backend.RoleGetListCommonReq) (res *backend.RoleGetListCommonRes, err error) {
	getListRes, err := service.Role().GetList(ctx, model.RoleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Total: getListRes.Total,
		Size:  getListRes.Size}, nil
}
