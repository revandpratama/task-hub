package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/task-hub/config"
	"github.com/revandpratama/task-hub/errorhandler"
	"github.com/revandpratama/task-hub/util"
)

func AuthorizationMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		token := ctx.Cookies("auth-token")
		if token == "" {
			return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: "unauthorized"})
		}

		isInvalidated, _ := config.RedisClient.SIsMember(context.Background(), "auth-token", token).Result()

		if isInvalidated {
			return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: "unauthorized"})
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
