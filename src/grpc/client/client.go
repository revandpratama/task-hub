package client

import (
	"log"

	"github.com/revandpratama/task-hub/grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient() user.UserServiceClient {
	conn, err := grpc.NewClient("microservice:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// (":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to port %v : %v", "9000", err)
	}

	auth := user.NewUserServiceClient(conn)

	return auth
	// response, err := auth.Login(context.Background(), &user.LoginRequest{Credential: "asdsa", Password: "asdasd"})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// log.Println("test")
	// log.Println(response.Token)
}
