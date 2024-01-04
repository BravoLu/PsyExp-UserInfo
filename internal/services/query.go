package services

import (
	"context"

	pb "github.com/BravoLu/grpc_idl"
	"github.com/grpc_plugins/log"
	"github.com/grpc_user_info/internal/mysql"
)

func (s *UserInfoServerImpl) QueryUser(
	ctx context.Context,
	req *pb.QueryReq,
) (*pb.QueryRsp, error) {
	log.Infof("req: %+v", req)
	dao := &mysql.UserInfoDaoImpl{}
	user, err := dao.QueryUser(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.QueryRsp{
		Code: 0,
		Msg: "ok",
		UserInfo: &pb.UserInfo{
			Uid: int64(user.ID),
			Email: user.Email,
			PhoneNumber: user.PhoneNumber,
			Gender: pb.GenderType(user.Gender),
			UserName: user.UserName,
			UserType: pb.UserType(user.UserType),
			Extra: user.Extra,
		},
	}, nil
}