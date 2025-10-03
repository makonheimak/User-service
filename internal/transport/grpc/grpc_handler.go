package grpc

import (
	"context"
	"log"

	pb "github.com/makonheimak/project-protos/proto/user"

	"github.com/makonheimak/user-service/internal/user/orm"
	"github.com/makonheimak/user-service/internal/user/service"

	"google.golang.org/protobuf/types/known/emptypb"
)

// grpcServer — это адаптер между protobuf и UserService
type Handler struct {
	pb.UnimplementedUserServiceServer
	svc *service.Service // ← указатель на сервис
}

// NewHandler создает новый обработчик
func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (s *Handler) PostUser(ctx context.Context, req *pb.PostUserRequest) (*pb.PostUserResponse, error) {
	log.Printf("gRPC PostUser called with email=%s", req.Email)

	user, err := s.svc.PostUser(orm.User{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &pb.PostUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Email: user.Email,
		},
	}, nil
}

// GetUserByID → вызывает GetUserByID
func (s *Handler) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := s.svc.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:    user.ID,
			Email: user.Email,
		},
	}, nil
}

// GetAllUsers → вызывает GetAllUsers
func (s *Handler) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	users, err := s.svc.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:    u.ID,
			Email: u.Email,
		})
	}

	return &pb.GetAllUsersResponse{Users: pbUsers}, nil
}

// PatchUserByID → вызывает PatchUserByID
func (s *Handler) PatchUserByID(ctx context.Context, req *pb.PatchUserByIDRequest) (*pb.PatchUserByIDResponse, error) {
	user, err := s.svc.PatchUserByID(req.Id, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.PatchUserByIDResponse{
		User: &pb.User{
			Id:    user.ID,
			Email: user.Email,
		},
	}, nil
}

// DeleteUserByID → вызывает DeleteUserByID
func (s *Handler) DeleteUserByID(ctx context.Context, req *pb.DeleteUserByIDRequest) (*emptypb.Empty, error) {
	err := s.svc.DeleteUserByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
