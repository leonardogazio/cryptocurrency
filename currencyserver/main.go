package main

import (
	"fmt"
	"log"
	"net"

	"github.com/leonardogazio/cryptocurrency/currencyserver/repo"
	"github.com/leonardogazio/cryptocurrency/currencyserver/server"
	"github.com/leonardogazio/cryptocurrency/proto/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Connecting to MongoDB...")
	if err := repo.Open(); err != nil {
		log.Fatalf("Could not connect to MongoDB: %s", err.Error())
	}

	// Start gRPC
	fmt.Println("Starting server on port :50051...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	fmt.Println("Server starting.")
	grpc := grpc.NewServer()
	pb.RegisterCurrencyServiceServer(grpc, &server.CurrencyServiceServer{})
	if err := grpc.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err.Error())
	}
}
