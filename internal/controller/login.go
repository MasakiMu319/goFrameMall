package controller

import (
	"context"
	"goFrameMall/api/backend"
	"goFrameMall/internal/model"
	"goFrameMall/internal/service"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

func (a *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
	res = &backend.LoginDoRes{}
	err = service.Login().Login(ctx, model.UserLoginInput{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	res.Info = service.Session().GetUser(ctx)
	return
}
