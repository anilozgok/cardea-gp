package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"os"
	"strings"
)

type DeletePhotoHandler struct {
	repo database.Repository
}

func NewDeletePhotoHandler(repo database.Repository) *DeletePhotoHandler {
	return &DeletePhotoHandler{repo: repo}
}

func (h *DeletePhotoHandler) Handle(c *fiber.Ctx) error {
	req := new(request.DeletePhotoRequest)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("error while parsing request")
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.repo.DeletePhotoById(c.Context(), req.PhotoID); err != nil {
		zap.L().Error("error while deleting photo", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("error while deleting photo")
	}

	photoPath := strings.TrimPrefix(req.PhotoURL, c.BaseURL())

	if err := os.Remove(strings.TrimPrefix(photoPath, "/")); err != nil {
		zap.L().Error("error while deleting photo file", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("error while deleting photo file")
	}

	return c.JSON(fiber.StatusOK)
}
