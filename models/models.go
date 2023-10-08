package models

import (
	pb "github.com/Joshua-FeatureFlag/proto/github.com/Joshua-FeatureFlag/proto/featureflag"
)

type Organization struct {
	ID    int64  `gorm:"primaryKey"`
	Users []User `gorm:"foreignKey:OrganizationID"`
	pb.Organization
}

type User struct {
	ID             int64 `gorm:"primaryKey"`
	OrganizationID int64
	pb.User
}
