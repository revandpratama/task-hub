package grpcserver

import (
	"fmt"
	"log"
	"net"
	"user-management-service/api/grpc/user"
	"user-management-service/internal/config"

	"google.golang.org/grpc"
)

func StartServer() {
	port := fmt.Sprintf(":%v", config.ENV.TCP_PORT)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to connect to tcp port 9000 : %v", err)
	}

	s := &user.Server{}

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server gRPC server to port 9000 : %v", err)
	}

	//client

}
