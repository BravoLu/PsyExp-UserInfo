package services

import (
	"context"

	pb "github.com/BravoLu/grpc_idl"
	"github.com/grpc_plugins/log"
	"github.com/grpc_user_info/internal/mysql"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserInfoServerImpl) Login(
	ctx context.Context,
	req *pb.LoginReq,
) (*pb.LoginRsp, error) {
	log.Infof("req: %+v", req)
	dao := &mysql.UserInfoDaoImpl{}
	 
	usr, err := dao.GetUser(ctx, req.Email)
	if err != nil {
		return &pb.LoginRsp{Code: 1, Msg: "email invalid", Uid: -1}, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(req.Password)); err != nil {
		return &pb.LoginRsp{Code: 2, Msg: "password, invalid", Uid: -1}, nil
	}

	return &pb.LoginRsp{Code: 0, Msg: "ok", Uid: int64(usr.ID)}, nil
}
