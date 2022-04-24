package logic

import (
	"context"

	"rpc/account"
	"rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExamTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExamTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExamTokenLogic {
	return &ExamTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExamTokenLogic) ExamToken(in *account.TokenExamineRequest) (*account.TokenExamineResponse, error) {
	// todo: add your logic here and delete this line
	in.GetToken()

	return &account.TokenExamineResponse{}, nil
}
