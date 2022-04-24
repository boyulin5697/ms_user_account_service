package logic

import (
	"context"

	"ms_user_account_service/api/internal/svc"
	"ms_user_account_service/api/internal/types"
	"ms_user_account_service/db/dao"
	"ms_user_account_service/db/entitydefines"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.GeneralResp, err error) {
	newUser := entitydefines.User{}
	newUser.SetName(req.Name)
	newUser.SetEmail(req.Email)
	newUser.SetPassword(req.Password)
	newUser.SetRight(1)
	newUser.SetTelephone(req.Telephone)
	newUser.SetNationality(req.Nationality)
	err = dao.AddUser(&newUser)
	if err != nil {
		var code int8 = 0
		resp.Code = code
		resp.Message = "Succeed!"
		resp.Information = ""
	} else {
		resp.Code = -1
		resp.Message = "Failed!"
		resp.Information = ""
	}
	return resp, err
}
