package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	pb "github.com/3FanYu/url-shortener-go/proto/main"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{
	*pb.UnimplementedShortUrlServer
}

func (s *server) CreateShortUrl(ctx context.Context, req *pb.CreateShortUrlReq) (*pb.CreateShortUrlResp, error) {
	// your implementation code here
	return &pb.CreateShortUrlResp{ShortUrl: "http://short.url/abc123"}, nil
}

func (s *server) RedirectShortUrl(ctx context.Context, req *pb.RedirectShortUrlReq) (*pb.RedirectShortUrlResp, error) {
	// your implementation code here
	return &pb.RedirectShortUrlResp{Url: "http://www.long.url/original/page"}, nil
}

func main() {
	// Create a listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register your gRPC service with the server
	pb.RegisterShortUrlServer(grpcServer, &server{})

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
	if err := pb.RegisterShortUrlHandler(ctx, mux, conn); err != nil {
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
