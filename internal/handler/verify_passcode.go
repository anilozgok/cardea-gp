package handler

import (
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

type VerifyPasscodeHandler struct {
}

func NewVerifyPasscodeHandler() *VerifyPasscodeHandler {
	return &VerifyPasscodeHandler{}
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

	// TODO:: line 33 throws nil pointer because c.Locals(passcodeCreatedAt) returns nil interface
	// TODO:: i think we cannot pass values from check-user's context to this request's context
	passcodeCreatedAt := c.Locals("passcodeCreatedAt").(time.Time)
	elapsed := time.Now().Sub(passcodeCreatedAt)

	if elapsed.Seconds() > 181.00 {
		zap.L().Info("TTL for verifying passcode has been expired.")
		clearContext(c, true)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	passcode := c.Locals("passcode").(int)
	if passcode != req.Passcode {
		zap.L().Info("incorrect passcode")
		clearContext(c, true)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	clearContext(c, false)

	return c.SendStatus(fiber.StatusOK)
}

func clearContext(c *fiber.Ctx, clearUser bool) {
	c.Locals("passcode", "")
	c.Locals("passcodeCreatedAt", "")

	if clearUser {
		c.Locals("user", "")
	}
}
