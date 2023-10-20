package api

import (
	"context"

	pb "github.com/Joshua-FeatureFlag/proto/github.com/Joshua-FeatureFlag/proto/system"
	"gorm.io/gorm"

	"github.com/Joshua-FeatureFlag/backend/api/organization"
	"github.com/Joshua-FeatureFlag/backend/api/user"
)

type Server struct {
	UserServer         pb.UserServiceServer
	OrganizationServer pb.OrganizationServiceServer
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		UserServer:         user.NewUserServiceServer(db),
		OrganizationServer: organization.NewOrganizationServiceServer(db),
	}
}

type UserService interface {
	CreateUser(ctx context.Context, req *pb.User) (*pb.User, error)
	GetUser(ctx context.Context, req *pb.User) (*pb.User, error)
	UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error)
	DeleteUser(ctx context.Context, req *pb.User) (*pb.User, error)
}

type OrganizationService interface {
	CreateOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error)
	GetOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error)
	UpdateOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error)
	DeleteOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error)
}
