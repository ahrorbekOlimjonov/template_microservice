package service

import (
	u "TEMPLATE_MICROSERVICE/genproto/user"
	"TEMPLATE_MICROSERVICE/pkg/logger"
	"TEMPLATE_MICROSERVICE/storage"
	"context"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	storage storage.IStorage
	Logger  logger.Logger
}

func NewUserService(db *sqlx.DB, log logger.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		Logger:  log,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *u.UserRequest) (*u.UserResponse, error) {
	res, err := s.storage.User().CreateUser(req)
	if err != nil {
		s.Logger.Error("Error insert user", logger.Any("Error insert user", err))
		return &u.UserResponse{}, status.Error(codes.Internal, "sothing went wrong, please chesk user info")
	}

	return res, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *u.UserId) (*u.UserResponse, error) {
	res, err := s.storage.User().GetUserById(req)
	if err != nil {
		s.Logger.Error("Error getting user", logger.Any("Error getting user", err))
		return &u.UserResponse{}, status.Error(codes.Internal, "sothing went wrong, please chesk user info")
	}

	return res, nil
}

func (s *UserService) GetUsersAll(ctx context.Context, req *u.UserListReq) (*u.Users, error) {
	res, err := s.storage.User().GetUsersAll(req)
	if err != nil {
		s.Logger.Error("Error getting all user", logger.Any("Error getting all user", err))
		return &u.Users{}, status.Error(codes.Internal, "sothing went wrong, please chesk user info")
	}

	return res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *u.UserId) (*u.Users, error) {
	res, err := s.storage.User().DeleteUser(req)

	if err != nil {
		s.Logger.Error("Error deleting user", logger.Any("Error deleting user", err))
		return &u.Users{}, status.Error(codes.Internal, "sothing went wrong, please chesk user info")
	}

	return res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *u.UserUpdateReq) (*u.UserResponse, error) {
	res, err := s.storage.User().UpdateUser(req)
	if err != nil {
		s.Logger.Error("Error updating user", logger.Any("Error updating user", err))
		return &u.UserResponse{}, status.Error(codes.Internal, "sothing went wrong, please chesk user info")
	}
	return res, nil
}

func (s *UserService) SearchUser(ctx context.Context, req *u.UserSearch) (*u.Users, error) {
	res, err := s.storage.User().SearchUser(req)
	if err != nil {
		s.Logger.Error("Error searching user", logger.Any("Error searching user", err))
		return &u.Users{}, status.Error(codes.Internal, "sothing went wrong, please chesk user info")
	}
	return res, nil
}
