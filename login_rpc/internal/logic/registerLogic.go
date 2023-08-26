package logic

import (
	"context"
	"errors"
	"login/login_rpc/model"

	"login/login_rpc/internal/svc"
	"login/login_rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	u, _ := l.svcCtx.UserModel.FindByUsername(l.ctx, in.Name)
	if u != nil {
		return nil, errors.New("用户名以存在")
	}

	user := &model.User{
		Name:     in.Name,
		Password: in.Password,
	}

	_, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResp{
		Flag: true,
	}, nil
}
