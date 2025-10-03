package grpc

import (
	"log"
	"net"

	"github.com/makonheimak/user-service/internal/user/service"

	pb "github.com/makonheimak/project-protos/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunGRPC запускает gRPC сервер
func RunGRPC(svc *service.Service) error {
	// 1. net.Listen на ":50052"
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}

	// 2. grpc.NewServer()
	grpcServer := grpc.NewServer()

	// 3. userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))
	pb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

	// 4. grpcSrv.Serve(listener)
	reflection.Register(grpcServer)

	log.Println("🚀 User Service gRPC server starting on :50052")
	return grpcServer.Serve(lis)
}
