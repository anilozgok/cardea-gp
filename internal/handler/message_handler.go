package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
)

type MessageHandler struct {
	repo database.Repository
}

func NewMessageHandler(repo database.Repository) *MessageHandler {
	return &MessageHandler{repo: repo}
}

func (h *MessageHandler) SendMessage(c *fiber.Ctx) error {
	message := new(entity.Message)
	if err := c.BodyParser(message); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	if err := h.repo.CreateMessage(c.Context(), message); err != nil {
		zap.L().Error("error while creating message", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(message)
}

func (h *MessageHandler) ListMessages(c *fiber.Ctx) error {
	userID1Str := c.Query("userID1")
	userID2Str := c.Query("userID2")

	if userID1Str == "" || userID2Str == "" {
		return c.Status(fiber.StatusBadRequest).SendString("userID1 and userID2 are required")
	}

	userID1, err := strconv.ParseUint(userID1Str, 10, 32)
	if err != nil {
		zap.L().Error("error while parsing userID1", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("Invalid userID1")
	}

	userID2, err := strconv.ParseUint(userID2Str, 10, 32)
	if err != nil {
		zap.L().Error("error while parsing userID2", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("Invalid userID2")
	}

	messages, err := h.repo.ListMessagesBetweenUsers(c.Context(), uint(userID1), uint(userID2))
	if err != nil {
		zap.L().Error("error while listing messages", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	return c.JSON(messages)
}
