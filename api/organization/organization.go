package organization

import (
	"context"

	pb "github.com/Joshua-FeatureFlag/proto/github.com/Joshua-FeatureFlag/proto/featureflag"
	"gorm.io/gorm"
)

type OrganizationServiceServer struct {
	pb.UnimplementedOrganizationServiceServer
	db *gorm.DB
}

func NewOrganizationServiceServer(db *gorm.DB) pb.OrganizationServiceServer {
	return &OrganizationServiceServer{db: db}
}

func (s OrganizationServiceServer) CreateOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error) {
	return nil, nil
}

func (s OrganizationServiceServer) GetOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error) {
	return nil, nil
}

func (s OrganizationServiceServer) UpdateOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error) {
	return nil, nil
}

func (s OrganizationServiceServer) DeleteOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error) {
	return nil, nil
}
