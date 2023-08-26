package logic

import (
	"context"
	"errors"
	"login/login_rpc/internal/svc"
	"login/login_rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {

	user, _ := l.svcCtx.UserModel.FindByUsername(l.ctx, in.Name)
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	if user.Password != in.Password {
		return nil, errors.New("密码错误")
	}

	return &pb.LoginResp{Flag: true}, nil
}
