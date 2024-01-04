package services

import (
	"context"
	pb "github.com/BravoLu/grpc_idl"
	"github.com/grpc_plugins/log"
	"github.com/grpc_user_info/internal/mysql"
)

func (s *UserInfoServerImpl) Exist(
	ctx context.Context,
	req *pb.ExistReq,
) (*pb.ExistRsp, error) {
	log.Infof("req: %+v", req)

	dao := &mysql.UserInfoDaoImpl{}
	usr, err := dao.GetUser(ctx, req.Email)
	if err != nil {
		return &pb.ExistRsp{Code: 1, Msg: "email invalid"}, nil
	}
	if usr == nil {
		return &pb.ExistRsp{Code: 1, Msg: "email invalid"}, nil
	}
	return &pb.ExistRsp{
		Code: 0,
		Msg: "success",
		Uid: int64(usr.ID),
	}, nil
}
