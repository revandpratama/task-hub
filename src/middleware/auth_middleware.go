package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/task-hub/errorhandler"
	"github.com/revandpratama/task-hub/util"
)

func AuthorizationMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		token := ctx.Cookies("auth-token")
		if token == "" {
			return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: "Unauthorized"})
		}

		userIDFromToken, userRoleFromToken, err := util.ValidateToken(token)
		if err != nil {
			return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: err.Error()})
		}

		ctx.Locals("userid", userIDFromToken)
		ctx.Locals("userrole", userRoleFromToken)

		return ctx.Next()
	}
}
