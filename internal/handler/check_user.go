package handler

import (
	"fmt"
	"github.com/anilozgok/cardea-gp/internal/config"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/pkg/mail"
	"github.com/anilozgok/cardea-gp/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type CheckUserHandler struct {
	repo   database.Repository
	config *config.Config
	fpCtx  *utils.ForgotPasswordCtx
}

func NewCheckUserHandler(repo database.Repository, config *config.Config, fpCtx *utils.ForgotPasswordCtx) *CheckUserHandler {
	return &CheckUserHandler{
		repo:   repo,
		config: config,
		fpCtx:  fpCtx,
	}
}

func (h *CheckUserHandler) Handle(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		zap.L().Error("email not found on the query param")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	maybeUser, err := h.repo.GetUserByEmail(c.Context(), email)
	if err != nil {
		zap.L().Error("error while checking user existence", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	if maybeUser == nil {
		zap.L().Error("user not found", zap.Error(err))
		return c.SendStatus(fiber.StatusNotFound)
	}

	passcode := rand.Intn(9000) + 1000

	mailServer := mail.NewMailServer(h.config.Secrets.EmailCredentials.Email, h.config.Secrets.EmailCredentials.Password)

	message := fmt.Sprintf("Your passcode is %d and it is valid for 3 minutes", passcode)

	now := time.Now()
	if err = mailServer.Send(email, message); err != nil {
		zap.L().Error("error while sending email.", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	h.fpCtx.Passcode = passcode
	h.fpCtx.CreatedAt = now
	h.fpCtx.User = *maybeUser

	return c.SendStatus(fiber.StatusOK)
}
