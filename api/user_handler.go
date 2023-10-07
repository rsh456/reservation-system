package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rsh456/reservation-system/db"
	"github.com/rsh456/reservation-system/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		LastName:  "Doe",
		FirstName: "Jhon",
	}
	return c.JSON(u)
}
