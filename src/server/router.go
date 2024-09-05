package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/revandpratama/task-hub/routes"
)

func InitRouters() *fiber.App {
	router := fiber.New()

	//define all the routes
	api := router.Group("/api")

	api.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))

	routes.AuthRoute(api)
	// routes.TaskRoutes(api)
	// routes.ProjectRoutes(api)
	routes.UserRoutes(api)

	return router
}
