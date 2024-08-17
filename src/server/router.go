package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/task-hub/routes"
)

func InitRouters() *fiber.App {
	router := fiber.New()

	//define all the routes
	api := router.Group("/api")

	routes.AuthRoute(api)
	routes.TaskRoutes(api)

	return router
}
