package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/anilozgok/cardea-gp/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UpdatePasswordHandler struct {
	repo  database.Repository
	fpCtx *utils.ForgotPasswordCtx
}

func NewUpdatePasswordHandler(repo database.Repository, fpCtx *utils.ForgotPasswordCtx) *UpdatePasswordHandler {
	return &UpdatePasswordHandler{
		repo:  repo,
		fpCtx: fpCtx,
	}
}

func (h *UpdatePasswordHandler) Handle(c *fiber.Ctx) error {
	/*if !h.fpCtx.Verified {
		zap.L().Info("passcode is not verified")
		return c.SendStatus(fiber.StatusBadRequest)
	}
	*/
	req := new(request.ForgotPassword)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		zap.L().Error("error while encrypting the password", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	if err = h.repo.UpdatePassword(c.Context(), string(hashedPassword), h.fpCtx.User); err != nil {
		zap.L().Error("error while updating the password", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
