//Login logic, Author @by created in 2022/4/27

package logic

import (
	"context"
	"strings"

	"ms_user_account_service/api/internal/svc"
	"ms_user_account_service/api/internal/types"

	"ms_user_account_service/db/dao"
	"ms_user_account_service/db/entitydefines"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.GeneralResp, err error) {
	var payload = req.LoginPayLoad
	var newUser *entitydefines.User
	res := types.GeneralResp{}
	if strings.Contains(payload, "@") == true{
		newUser = dao.GetUserByEmail(&payload)
		if nil == newUser{
			res.Code = -1
			res.Message = "Login failed!"
			res.Information = "The user is not exist!"
			return &res,err
		}
		passedPass := newUser.Password()
		if passedPass==req.Password {
			return loginSuccess(&res),nil
		}else{
			res.Code = -2
			res.Message = "Login failed!"
			res.Information = "The password is wrong!"
			return &res, nil
		}
	}else{
		newUser = dao.GetUserByName(&payload)
		if nil == newUser{
			res.Code = -1
			res.Message = "Login failed!"
			res.Information = "The user is not exist!"
			return &res,err
		}
		passedPass := newUser.Password()
		if passedPass == req.Password{
			return loginSuccess(&res),nil
		}else{
			res.Code = -2
			res.Message = "Login failed!"
			res.Information = "The password is wrong!"
			return &res, nil
		}
	}
}

func loginSuccess(resp *types.GeneralResp) *types.GeneralResp{
	resp.Code = 0
	resp.Message = "Login success!"
	return resp
}


