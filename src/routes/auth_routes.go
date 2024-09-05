package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/task-hub/grpc/client"
	"github.com/revandpratama/task-hub/handler"
	"github.com/revandpratama/task-hub/middleware"
)

func AuthRoute(r fiber.Router) {

	auth := client.NewAuthClient()
	handler := handler.NewAuthHandler(auth)

	routeGuest := r.Group("/auth")

	routeGuest.Use(middleware.GuestMiddleware())

	routeGuest.Post("/login", handler.Login)

	routeGuest.Post("/register", handler.Register)

}
