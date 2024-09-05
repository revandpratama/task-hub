package main

import (
	grpcserver "user-management-service/api/grpc/grpc-server"
	"user-management-service/internal/config"
	"user-management-service/internal/database"
)

func main() {
	config.LoadConfig()
	database.LoadDB()

	grpcserver.StartServer()

}

// func InitService() {

// }
