package logic

import (
	"context"
	"login/login_rpc/pb"

	"login/login_api/internal/svc"
	"login/login_api/internal/types"

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

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.UserRpcClient.Register(l.ctx, &pb.RegisterReq{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{
		Flag: registerResp.Flag,
	}, nil
}
