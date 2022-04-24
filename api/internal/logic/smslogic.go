package logic

import (
	"context"

	"ms_user_account_service/api/internal/svc"
	"ms_user_account_service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SMSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SMSLogic {
	return &SMSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SMSLogic) SMS() (resp *types.GeneralResp, err error) {
	// todo: add your logic here and delete this line

	return
}
