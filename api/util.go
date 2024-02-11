package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rsh456/reservation-system/types"
)

func getAuthUser(c *fiber.Ctx) (*types.User, error) {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return user, nil
}
