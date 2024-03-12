package handler

import (
	"net/http"
	"strconv"
	"test-grpc-project/pkg/model"
	"test-grpc-project/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
}

type userHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) IUserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Create(ctx *fiber.Ctx) error {
	var user model.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	if err := h.service.CreateUser(&user); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"success": "user created successfully",
		"user":    user,
	})
}

func (h *userHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idToi, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			map[string]interface{}{
				"error": err.Error(),
			},
		)
	}

	user, err := h.service.GetUser(uint32(idToi))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			map[string]interface{}{
				"error": "user not found!",
			},
		)
	}

	if err := h.service.DeleteUser(uint32(idToi)); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			map[string]interface{}{
				"error": "user deletion failed",
			},
		)
	}

	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"success": "user deleted successfully",
		"user":    user,
	})
}

func (h *userHandler) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idToi, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			map[string]interface{}{
				"error": err.Error(),
			},
		)
	}
	user, err := h.service.GetUser(uint32(idToi))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			map[string]interface{}{
				"error": "user not found!",
			},
		)
	}

	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"success": "user getting is successful",
		"user":    user,
	})
}
