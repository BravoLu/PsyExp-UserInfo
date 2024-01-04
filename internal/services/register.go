package services

import (
	"context"

	pb "github.com/BravoLu/grpc_idl"
	"github.com/grpc_plugins/log"
	"github.com/grpc_user_info/internal/entity"
	"github.com/grpc_user_info/internal/mysql"
)

func (u *UserInfoServerImpl) Register(
	ctx context.Context,
	req *pb.RegisterReq,
) (*pb.RegisterRsp, error) {
	log.Infof("req: %+v", req)
	user := &entity.UserInfo{
		Email:       req.UserInfo.Email,
		UserName:    req.UserInfo.UserName,
		PhoneNumber: req.UserInfo.PhoneNumber,
		Gender:      uint32(req.UserInfo.Gender),
		UserType:    uint32(req.UserInfo.UserType),
		Password:    req.Password,
	}

	dao := &mysql.UserInfoDaoImpl{}
	_, err := dao.AddUser(ctx, user)
	if err != nil {
		return &pb.RegisterRsp{Code: 10, Msg: "The email has been registered or data is invalid."}, nil
	}
	return &pb.RegisterRsp{Code: 0, Msg: "ok"}, nil
}
