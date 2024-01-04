package services

import (
	pb "github.com/BravoLu/grpc_idl"
)

type UserInfoServerImpl struct {
	pb.UnimplementedUserServiceServer
}
