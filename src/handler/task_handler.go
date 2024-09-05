package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/task-hub/dto"
	"github.com/revandpratama/task-hub/errorhandler"
	"github.com/revandpratama/task-hub/service"
	"github.com/revandpratama/task-hub/util"
)

type taskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *taskHandler {
	return &taskHandler{
		service: service,
	}
}

func (h taskHandler) GetAll(ctx *fiber.Ctx) error {
	tasks, err := h.service.GetAll()
	if err != nil {
		// return errorhandler.HandleError(ctx, err)
		return errorhandler.HandleError(ctx, &errorhandler.NotFoundErr{Message: "task empty"})
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve all tasks",
		Data:       tasks,
	})

	return ctx.JSON(response)
}

func (h taskHandler) GetAllUserTask(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("userid"))

	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: "bad request parameter"})
	}

	userIDFromToken := ctx.Locals("userid").(*int)
	userRoleFromToken := ctx.Locals("userrole").(*string)
	if userID != *userIDFromToken && *userRoleFromToken != "admin" {
		return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: "unauthorized"})
	}

	tasks, err := h.service.GetByUserID(userID)
	if err != nil {
		return errorhandler.HandleError(ctx, err)
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "retrieve all task by user id success",
		Data:       tasks,
	})

	return ctx.JSON(response)
}

func (h taskHandler) GetUserTaskByID(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: "bad request parameter"})
	}

	userIDFromToken := ctx.Locals("userid").(*int)
	userRoleFromToken := ctx.Locals("userrole").(*string)
	if userID != *userIDFromToken && *userRoleFromToken != "admin" {
		return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: "unauthorized"})
	}

	taskID, err := strconv.Atoi(ctx.Params("taskid"))
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: "bad request parameter"})
	}

	task, err := h.service.Get(taskID)
	if err != nil {
		return errorhandler.HandleError(ctx, err)
	}

	if task.UserID != *userIDFromToken && *userRoleFromToken != "admin" {
		return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: "unauthorized"})
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "retrieve one task by user id and task id success",
		Data:       task,
	})

	return ctx.JSON(response)
}

func (h taskHandler) Get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
	}

	task, err := h.service.Get(id)
	if err != nil {
		return errorhandler.HandleError(ctx, err)
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve task",
		Data:       task,
	})

	return ctx.JSON(response)

}

func (h taskHandler) Create(ctx *fiber.Ctx) error {
	var request dto.TaskRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
	}

	if err := h.service.Create(request); err != nil {
		return errorhandler.HandleError(ctx, err)
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "success create task",
	})

	return ctx.JSON(response)
}

func (h taskHandler) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
	}

	var request dto.TaskRequest
	if err := ctx.BodyParser(&request); err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
	}

	if err := h.service.Update(id, request); err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.InternalServerErr{Message: err.Error()})
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "success update task",
	})

	return ctx.JSON(response)
}

func (h taskHandler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
	}

	if err := h.service.Delete(id); err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.InternalServerErr{Message: err.Error()})
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "success delete task",
	})

	return ctx.JSON(response)

}
