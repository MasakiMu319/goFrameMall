package controller

import (
	"context"
	"goFrameMall/api/backend"
	"goFrameMall/internal/service"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

//func (a *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
//	res = &backend.LoginDoRes{}
//	err = service.Login().Login(ctx, model.UserLoginInput{
//		Name:     req.Name,
//		Password: req.Password,
//	})
//	if err != nil {
//		return
//	}
//	//res.Info = service.Session().GetUser(ctx)
//	return
//}

func (c *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
	res = &backend.LoginDoRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	return
}

func (c *cLogin) RefreshToken(ctx context.Context, req *backend.LoginRefreshTokenReq) (res *backend.LoginRefreshTokenRes, err error) {
	res = &backend.LoginRefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

func (c *cLogin) Logout(ctx context.Context, req *backend.LoginLogoutReq) (res *backend.LoginLogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
