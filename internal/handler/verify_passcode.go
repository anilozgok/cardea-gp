package handler

import (
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/anilozgok/cardea-gp/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

type VerifyPasscodeHandler struct {
	fpCtx *utils.ForgotPasswordCtx
}

func NewVerifyPasscodeHandler(fpCtx *utils.ForgotPasswordCtx) *VerifyPasscodeHandler {
	return &VerifyPasscodeHandler{
		fpCtx: fpCtx,
	}
}

func (h *VerifyPasscodeHandler) Handle(c *fiber.Ctx) error {
	req := new(request.VerifyPasscode)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return err
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return err
	}

	elapsed := time.Now().Sub(h.fpCtx.CreatedAt)

	if elapsed.Seconds() > 181.00 {
		h.fpCtx.Expired = true
		zap.L().Info("TTL for verifying passcode has been expired.")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if h.fpCtx.Passcode != req.Passcode {
		zap.L().Info("incorrect passcode")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}
