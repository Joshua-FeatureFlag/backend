package user

import (
	"context"

	pb "github.com/Joshua-FeatureFlag/proto/github.com/Joshua-FeatureFlag/proto/featureflag"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	db *gorm.DB
}

func NewUserServiceServer(db *gorm.DB) pb.UserServiceServer {
	return &UserServiceServer{db: db}
}

func (s UserServiceServer) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return nil, nil
}

func (s UserServiceServer) GetUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return nil, nil
}

func (s UserServiceServer) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return nil, nil
}

func (s UserServiceServer) DeleteUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return nil, nil
}
