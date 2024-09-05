package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/task-hub/config"
	"github.com/revandpratama/task-hub/database"
	"github.com/revandpratama/task-hub/handler"
	"github.com/revandpratama/task-hub/middleware"
	"github.com/revandpratama/task-hub/repository"
	"github.com/revandpratama/task-hub/service"
)

func UserRoutes(r fiber.Router) {
	taskRepo := repository.NewTaskRepository(database.DBCONN, config.RedisClient)
	taskAttRepo := repository.NewTaskAttachmentRepository(database.DBCONN)
	projectRepo := repository.NewProjectRepository(database.DBCONN)

	taskService := service.NewTaskService(taskRepo, taskAttRepo)
	projectService := service.NewProjectService(projectRepo, taskRepo)

	taskHandler := handler.NewTaskHandler(taskService)
	projectHandler := handler.NewProjectHandler(projectService)

	ur := r.Group("/:userid")

	ur.Use(middleware.AuthorizationMiddleware())
	
	ur.Get("/tasks", taskHandler.GetAllUserTask)
	ur.Get("/tasks/:taskid", taskHandler.GetUserTaskByID)
	ur.Post("/tasks/:taskid", taskHandler.Create)
	ur.Put("/tasks/:taskid", taskHandler.Update)
	ur.Delete("/tasks/:taskid", taskHandler.Delete)

	ur.Get("/projects", projectHandler.GetAllUserProject)
	ur.Get("/projects/:projectid", projectHandler.GetUserProjectByID)
	ur.Post("/projects/:projectid", projectHandler.Create)
	ur.Put("/projects/:projectid", projectHandler.Update)
	ur.Delete("/projects/:projectid", projectHandler.Delete)

}
