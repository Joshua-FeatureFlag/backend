package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Joshua-FeatureFlag/proto/github.com/Joshua-FeatureFlag/proto/featureflag"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Joshua-FeatureFlag/backend/api"
	"github.com/Joshua-FeatureFlag/backend/migrate"
)

const (
	address = "localhost:50051"
	dsn     = "user=featureflag password=password dbname=featureflag_dev host=host.docker.internal port=5432 sslmode=disable"
)

func main() {
	action := flag.String("action", "serve", "Action to perform: 'migrate' or 'serve'")
	flag.Parse()

	// Initialize database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	switch *action {
	case "migrate":
		migrate.RunMigration(db)
		if err != nil {
			log.Fatalf("Failed to migrate database schema: %v", err)
		}
		fmt.Println("Database migration completed successfully")
	case "serve":
		// Create new api.Server instance
		server := api.NewServer(db)

		// Start gRPC server
		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterUserServiceServer(grpcServer, server.UserServer)
		pb.RegisterOrganizationServiceServer(grpcServer, server.OrganizationServer)

		fmt.Printf("gRPC server listening on %s\n", address)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	default:
		log.Fatalf("Unknown action: %s. Supported actions are 'migrate' or 'serve'", *action)
	}
}
