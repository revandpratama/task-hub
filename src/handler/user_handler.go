package handler

import (
	"github.com/revandpratama/task-hub/service"
)

type userHandler struct {
	taskService    service.TaskService
	projectService service.ProjectService
}

func NewUserHandler(taskService service.TaskService, projectService service.ProjectService) *userHandler {
	return &userHandler{
		taskService:    taskService,
		projectService: projectService,
	}
}
