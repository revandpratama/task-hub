package errorhandler

// import (
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/revandpratama/task-hub/dto"
// 	"github.com/revandpratama/task-hub/util"
// )

// func HandleError(c *fiber.Ctx, err error) error {
// 	var statusCode int

// 	switch err.(type) {
// 	case *NotFoundErr:
// 		statusCode = http.StatusNotFound
// 	case *InternalServerErr:
// 		statusCode = http.StatusInternalServerError
// 	case *UnauthorizedErr:
// 		statusCode = http.StatusUnauthorized
// 	case *BadRequestErr:
// 		statusCode = http.StatusBadRequest
// 	}

// 	response := util.NewResponse(dto.ResponseParam{
// 		StatusCode: statusCode,
// 		Message:    err.Error(),
// 	})

// 	return c.JSON(response)

// }
