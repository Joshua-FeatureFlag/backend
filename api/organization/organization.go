package organization

import (
	"context"
	"fmt"

	pb "github.com/Joshua-FeatureFlag/proto/gen/gen/go"
	"gorm.io/gorm"
)

type OrganizationServiceServer struct {
	pb.UnimplementedOrganizationServiceServer
	db *gorm.DB
}

func NewOrganizationServiceServer(db *gorm.DB) pb.OrganizationServiceServer {
	return &OrganizationServiceServer{db: db}
}

func (s *OrganizationServiceServer) CreateOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error) {
	organization := &pb.Organization{
		Name: req.GetName(),
	}
	result := s.db.Create(organization)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create organization: %w", result.Error)
	}
	return organization, nil
}

func (s *OrganizationServiceServer) GetOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error) {
	organization := &pb.Organization{}
	result := s.db.First(organization, req.GetId())
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get organization: %w", result.Error)
	}
	return organization, nil
}

func (s *OrganizationServiceServer) UpdateOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error) {
	organization := &pb.Organization{}
	result := s.db.First(organization, req.GetId())
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find organization: %w", result.Error)
	}
	organization.Name = req.GetName()
	result = s.db.Save(organization)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update organization: %w", result.Error)
	}
	return organization, nil
}

func (s *OrganizationServiceServer) DeleteOrganization(ctx context.Context, req *pb.Organization) (*pb.Organization, error) {
	organization := &pb.Organization{}
	result := s.db.Delete(organization, req.GetId())
	if result.Error != nil {
		return nil, fmt.Errorf("failed to delete organization: %w", result.Error)
	}
	return organization, nil
}
