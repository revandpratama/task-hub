package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/task-hub/dto"
	"github.com/revandpratama/task-hub/errorhandler"
	"github.com/revandpratama/task-hub/service"
	"github.com/revandpratama/task-hub/util"
)

type projectHandler struct {
	service service.ProjectService
}

func NewProjectHandler(service service.ProjectService) *projectHandler {
	return &projectHandler{
		service: service,
	}
}

func (h projectHandler) GetAll(ctx *fiber.Ctx) error {
	projects, err := h.service.GetAll()

	if err != nil || projects == nil {
		return errorhandler.HandleError(ctx, err)
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve all projects",
		Data:       projects,
	})

	return ctx.JSON(response)

}

func (h projectHandler) GetById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
	}

	project, err := h.service.GetById(id)
	if err != nil {
		return errorhandler.HandleError(ctx, err)
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve project by id",
		Data:       project,
	})

	return ctx.JSON(response)
}

func (h projectHandler) Create(ctx *fiber.Ctx) error {
	var request dto.ProjectRequest
	if err := ctx.BodyParser(&request); err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
	}

	err := h.service.Create(request)
	if err != nil {
		return errorhandler.HandleError(ctx, err)
	}

	respose := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "success create project",
	})

	return ctx.JSON(respose)
}

func (h projectHandler) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
	}

	var request dto.ProjectRequest
	if err := ctx.BodyParser(&request); err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
	}

	if err := h.service.Update(id, request); err != nil {
		return errorhandler.HandleError(ctx, err)
	}

	respose := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "success update project",
	})

	return ctx.JSON(respose)
}

func (h projectHandler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
	}

	if err := h.service.Delete(id); err != nil {
		return errorhandler.HandleError(ctx, err)
	}

	respose := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "success delete project",
	})

	return ctx.JSON(respose)
}

// func (h projectHandler) Update(ctx *fiber.Ctx) error {
// 	id, err := ctx.ParamsInt("id")
// 	if err != nil {
// 		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
// 	}

// 	var request dto.ProjectRequest
// 	if err := ctx.BodyParser(&request); err != nil {
// 		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: err.Error()})
// 	}

// 	err := h.service.Delete(id)
// 	if err != nil {
// 		return errorhandler.HandleError(ctx, err)
// 	}

// 	respose := util.NewResponse(dto.ResponseParam{
// 		StatusCode: fiber.StatusOK,
// 		Message:    "success create project",
// 	})

// 	return ctx.JSON(respose)
// }

func (h projectHandler) GetAllUserProject(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: "bad request parameter"})
	}

	userIDFromToken := ctx.Locals("userid").(*int)
	userRoleFromToken := ctx.Locals("userrole").(*string)
	if userID != *userIDFromToken && *userRoleFromToken != "admin" {
		return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: "unauthorized"})
	}

	projects, err := h.service.GetAllUserProject(userID)
	if err != nil {
		return errorhandler.HandleError(ctx, err)
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "retrieve all user project success",
		Data:       projects,
	})

	return ctx.JSON(response)
}

func (h projectHandler) GetUserProjectByID(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: "bad request parameter"})
	}

	userIDFromToken := ctx.Locals("userid").(*int)
	userRoleFromToken := ctx.Locals("userrole").(*string)
	if userID != *userIDFromToken && *userRoleFromToken != "admin" {
		return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: "unauthorized"})
	}

	projectID, err := strconv.Atoi(ctx.Params("projectid"))
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.BadRequestErr{Message: "bad request parameter"})
	}

	project, err := h.service.GetById(projectID)
	if err != nil {
		return errorhandler.HandleError(ctx, &errorhandler.InternalServerErr{Message: err.Error()})
	}

	if project.UserID != *userIDFromToken && *userRoleFromToken != "admin" {
		return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: "unauthorized"})
	}

	response := util.NewResponse(dto.ResponseParam{
		StatusCode: fiber.StatusOK,
		Message:    "retrieve one project by user id and project id success",
		Data:       project,
	})

	return ctx.JSON(response)
}
