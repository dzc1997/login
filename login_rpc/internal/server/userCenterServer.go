// Code generated by goctl. DO NOT EDIT.
// Source: login.proto

package server

import (
	"context"

	"login/login_rpc/internal/logic"
	"login/login_rpc/internal/svc"
	"login/login_rpc/pb"
)

type UserCenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserCenterServer
}

func NewUserCenterServer(svcCtx *svc.ServiceContext) *UserCenterServer {
	return &UserCenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UserCenterServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserCenterServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}