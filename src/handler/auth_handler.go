package handler

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/task-hub/config"
	"github.com/revandpratama/task-hub/dto"
	"github.com/revandpratama/task-hub/errorhandler"
	"github.com/revandpratama/task-hub/grpc/user"
	"github.com/revandpratama/task-hub/util"
)

type authHandler struct {
	auth user.UserServiceClient
}

func NewAuthHandler(auth user.UserServiceClient) *authHandler {
	return &authHandler{
		auth: auth,
	}
}

func (h authHandler) Login(ctx *fiber.Ctx) error {

	var request dto.LoginRequest

	if err := ctx.BodyParser(&request); err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: "wrong credentials"})
	}

	creds := user.LoginRequest{
		Credential: request.Identifier,
		Password:   request.Password,
	}

	res, err := h.auth.Login(context.Background(), &creds)
	if err != nil || res.Error != "" {

		response := util.NewResponse(dto.ResponseParam{
			StatusCode: fiber.StatusOK,
			Message:    fmt.Sprintf("login failed: %v", res.Error),
		})

		return ctx.JSON(response)

	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "login success",
		Data:       dto.LoginResponse{AccessToken: res.Token},
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "auth-token",
		Value:    res.Token,
		Domain:   config.ENV.DOMAIN,
		Path:     "/",
		HTTPOnly: false,
		Secure:   false,
		MaxAge:   100,
	})

	return ctx.JSON(response)
}

func (h authHandler) Register(ctx *fiber.Ctx) error {

	var request dto.RegisterRequest

	if err := ctx.BodyParser(&request); err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: "wrong credentials"})
	}

	newUser := user.RegisterRequest{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	err, error := h.auth.Register(context.Background(), &newUser)

	if err != nil || error != nil {
		return error
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "register success",
	})

	return ctx.JSON(response)
}
