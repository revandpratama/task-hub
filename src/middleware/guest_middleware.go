package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/task-hub/errorhandler"
)

func GuestMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		token := ctx.Cookies("auth-token")
		if token != "" {
			return errorhandler.HandleError(ctx, &errorhandler.UnauthorizedErr{Message: "guest only route"})
		}

		return ctx.Next()
	}
}
