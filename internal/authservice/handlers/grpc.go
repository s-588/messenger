package handlers

import (
	"github.com/s-588/messenger/internal/authservice/service"
	pb "github.com/s-588/messenger/internal/genproto/authservice"
)

type GRPCServer struct  {
	pb.UnimplementedAuthServiceServer
	svc *service.MessageService
}

func NewGRPCServer(svc *service.MessageService) *GRPCServer  {
	return &GRPCServer{
		svc: svc,
	}
}
