package services

import (
	"context"

	pb "github.com/BravoLu/grpc_idl"
	"github.com/grpc_plugins/log"
	"github.com/grpc_user_info/internal/mysql"
)

func (s *UserInfoServerImpl) QueryUsers(
	ctx context.Context,
	req *pb.QueryUsersReq,
) (*pb.QueryUsersRsp, error) {
	log.Infof("req: %+v", req)

	dao := &mysql.UserInfoDaoImpl{}
	var uids []int64
	for _, u := range req.Uids {
		uids = append(uids, u)
	}
	us, err := dao.QueryUsers(ctx, uids)

	if err != nil {
		return &pb.QueryUsersRsp{Code: 5, Msg: "invalid uids"}, nil
	}
	var pbUser []*pb.UserInfo
	for _, v := range us {
		pbUser = append(pbUser, &pb.UserInfo{
			Uid:      int64(v.ID),
			Email:    v.Email,
			UserName: v.UserName,
		})
	}

	return &pb.QueryUsersRsp{Code: 0, Msg: "success", Users: pbUser}, nil
}
