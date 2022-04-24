package logic

import (
	"context"

	"ms_user_account_service/api/internal/svc"
	"ms_user_account_service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailLogic {
	return &EmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailLogic) Email() (resp *types.GeneralResp, err error) {
	// todo: add your logic here and delete this line

	return
}
