package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/Joshua-FeatureFlag/proto/gen/gen/go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Joshua-FeatureFlag/backend/api"
	"github.com/Joshua-FeatureFlag/backend/middleware"
	"github.com/Joshua-FeatureFlag/backend/migrate"
)

const (
	grpc_address = ":50051"

	dsn = "user=featureflag password=password dbname=featureflag_dev host=host.docker.internal port=5432 sslmode=disable"
)

func main() {
	action := flag.String("action", "serve", "Action to perform: 'migrate' or 'serve'")
	flag.Parse()

	switch *action {
	case "migrate":
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		migrate.RunMigration(db)
		if err != nil {
			log.Fatalf("Failed to migrate database schema: %v", err)
		}
		fmt.Println("Database migration completed successfully")
	case "serve":
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Create new api.Server instance
		server := api.NewServer(db)

		// Start gRPC server
		lis, err := net.Listen("tcp", grpc_address)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterUserServiceServer(grpcServer, server.UserServer)
		pb.RegisterOrganizationServiceServer(grpcServer, server.OrganizationServer)

		fmt.Printf("gRPC server listening on %s\n", grpc_address)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	case "http":

		// Load environment variables from a .env file
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading the .env file: %v", err)
		}

		ctx := context.Background()
		mux := runtime.NewServeMux()
		customMux := &middleware.CustomMux{
			Mux: mux,
			Endpoint: []middleware.Endpoint{
				{
					Path: "/v1/user/",
					Middleware: []middleware.Middleware{
						middleware.EnsureValidToken(),
						middleware.EnsureValidScope("user:read"),
					},
					Handler: mux,
				},
			},
		}
		opts := []grpc.DialOption{grpc.WithInsecure()}

		err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpc_address, opts)
		if err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
		err = pb.RegisterOrganizationServiceHandlerFromEndpoint(ctx, mux, grpc_address, opts)
		if err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}

		corsWrapper := middleware.CORSMiddleware(customMux)

		// Start HTTP server
		log.Print("Server listening on :5000")
		err = http.ListenAndServe(":5000", corsWrapper)
		if err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}

	default:
		log.Fatalf("Unknown action: %s. Supported actions are 'migrate' or 'serve' or 'http'", *action)
	}
}
