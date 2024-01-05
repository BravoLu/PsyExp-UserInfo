package services

import (
	"context"

	pb "github.com/BravoLu/grpc_idl"
	"github.com/grpc_plugins/log"
	"github.com/grpc_user_info/internal/entity"
	"github.com/grpc_user_info/internal/mysql"
)

func (u *UserInfoServerImpl) UpdateUser(
	ctx context.Context,
	req *pb.UpdateUserReq,
) (*pb.UpdateUserRsp, error) {
	log.Infof("req: %+v", req)
	user := &entity.UserInfo{
		Password: req.Password,
		Extra: req.Extra,
		PhoneNumber: req.PhoneNumber,
	}

	dao := &mysql.UserInfoDaoImpl{}
	err := dao.UpdateUser(ctx, req.Uid, user)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserRsp{Code: 0, Msg: "ok"}, nil
}
