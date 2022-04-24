package logic

import (
	"context"

	"ms_user_account_service/api/internal/svc"
	"ms_user_account_service/api/internal/types"

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

func (l *LoginLogic) Login(req *types.GeneralResp) (resp *types.GeneralResp, err error) {
	// todo: add your logic here and delete this line

	return
}
