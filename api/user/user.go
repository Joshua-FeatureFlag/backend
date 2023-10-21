package user

import (
	"context"
	"fmt"

	pb "github.com/Joshua-FeatureFlag/proto/gen/gen/go"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	db *gorm.DB
}

func NewUserServiceServer(db *gorm.DB) pb.UserServiceServer {
	return &UserServiceServer{db: db}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := &pb.User{
		Name:           req.GetName(),
		OrganizationId: req.GetOrganizationId(),
	}
	result := s.db.Create(user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create user: %w", result.Error)
	}
	return user, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := &pb.User{}
	result := s.db.First(user, req.GetId())
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user: %w", result.Error)
	}
	return user, nil
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := &pb.User{}
	result := s.db.First(user, req.GetId())
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find user: %w", result.Error)
	}
	user.Name = req.GetName()
	result = s.db.Save(user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update user: %w", result.Error)
	}
	return user, nil
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := &pb.User{}
	result := s.db.Delete(user, req.GetId())
	if result.Error != nil {
		return nil, fmt.Errorf("failed to delete user: %w", result.Error)
	}
	return user, nil
}
