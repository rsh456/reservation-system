package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rsh456/reservation-system/types"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return fmt.Errorf("not authorized")
	}
	if !user.IsAdmin {
		return fmt.Errorf("not authorized")
	}
	return c.Next()
}
