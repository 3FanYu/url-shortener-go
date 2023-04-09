package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	urlShortenerRepo "github.com/3FanYu/url-shortener-go/adapter/repository/mongo/url_shortener"
	pb "github.com/3FanYu/url-shortener-go/proto/short_url"
	urlShortenerHlr "github.com/3FanYu/url-shortener-go/router/handler/url_shortener"
	urlShortenerUC "github.com/3FanYu/url-shortener-go/usecase/url_shortener"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://db:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	db := client.Database("url_shortener")

	repo := urlShortenerRepo.NewRepository(db)
	uc := urlShortenerUC.NewUsecase(repo)
	hlr := urlShortenerHlr.NewHandler(uc)

	// Create a listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register your gRPC service with the server
	pb.RegisterUrlShortenerServer(grpcServer, hlr)

	// Register gRPC reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Start the gRPC server in a goroutine
	go func() {
		fmt.Println("Starting gRPC server on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Printf("Failed to start gRPC server: %v\n", err)
			return
		}
	}()

	// Create a context with a cancel function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a gRPC dial option with insecure transport
	dialOption := grpc.WithInsecure()

	// Dial the gRPC server
	conn, err := grpc.DialContext(ctx, "localhost:50051", dialOption)
	if err != nil {
		fmt.Printf("Failed to dial server: %v\n", err)
		return
	}
	defer conn.Close()

	// Create a new Mux
	mux := runtime.NewServeMux()

	// Register your gRPC service on the Mux
	if err := pb.RegisterUrlShortenerHandler(ctx, mux, conn); err != nil {
		fmt.Printf("Failed to register service: %v\n", err)
		return
	}

	// Start the HTTP server
	fmt.Println("Starting HTTP server on :3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
}
