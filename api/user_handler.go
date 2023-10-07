package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rsh456/reservation-system/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		LastName:  "Doe",
		FirstName: "Jhon",
	}
	return c.JSON(u)
}
func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Doe")
}
