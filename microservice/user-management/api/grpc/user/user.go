package user

import (
	"context"
	"user-management-service/api/dto"
	"user-management-service/internal/database"
	"user-management-service/internal/repository"
	"user-management-service/internal/usecase"
)

type Server struct {
	UnimplementedUserServiceServer
}

func (s *Server) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	userRepo := repository.NewUserRepository(database.DBCONN)
	userService := usecase.NewUserService(userRepo)

	requestToService := dto.LoginRequest{
		Credential: request.Credential,
		Password:   request.Password,
	}

	tokenString, err := userService.Login(requestToService)
	if err != nil {
		return &LoginResponse{
			Error: err.Error(),
		}, nil
	}
	return &LoginResponse{
		Token: *tokenString,
	}, nil

	// return response, nil
}

func (s *Server) Register(ctx context.Context, request *RegisterRequest) (*ErrorResponse, error) {
	userRepo := repository.NewUserRepository(database.DBCONN)
	userService := usecase.NewUserService(userRepo)

	requestToService := dto.RegisterRequest{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := userService.Register(requestToService); err != nil {
		return &ErrorResponse{
			Error: err.Error(),
		}, nil
	}

	return nil, nil
}
