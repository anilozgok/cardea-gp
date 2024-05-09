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

	emailTemplate := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>OTP Email</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;">
			<table cellpadding="0" cellspacing="0" width="100%%" style="max-width: 600px; margin: 0 auto; background-color: #fff; border-radius: 8px; box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1); padding: 20px;">
				<tr>
					<td align="center" style="padding: 20px 0;">
						<h2 style="margin-bottom: 20px;">Your OTP Code</h2>
						<p style="margin-bottom: 20px;">Use the following OTP code to complete your verification:</p>
						<div style="background-color: #007bff; color: #fff; padding: 10px 20px; border-radius: 4px; font-size: 24px; margin-bottom: 20px;">%d</div>
						<p style="margin-bottom: 20px;">This OTP code is valid for a limited time. Do not share it with anyone.</p>
					</td>
				</tr>
			</table>
		</body>
		</html>
	`
	passcode := rand.Intn(9000) + 1000

	mailServer := mail.NewMailServer(h.config.Secrets.EmailCredentials.Email, h.config.Secrets.EmailCredentials.Password)

	message := fmt.Sprintf(emailTemplate, passcode)

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
