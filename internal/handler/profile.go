package handler

import (
	"fmt"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ProfileHandler struct {
	repo database.Repository
}

func NewProfileHandler(repo database.Repository) *ProfileHandler {
	return &ProfileHandler{repo: repo}
}

func (h *ProfileHandler) CreateProfile(c *fiber.Ctx) error {
	req := new(request.CreateProfileRequest)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("Validation failed: " + err.Error())
	}

	profile := &entity.Profile{
		UserID:            req.UserID,
		Bio:               req.Bio,
		ProfilePictureURL: req.ProfilePicture,
		Experience:        req.Experience,
		Specialization:    req.Specialization,
		Photos:            req.Photos,
	}

	if err := h.repo.CreateProfile(c.Context(), profile); err != nil {
		zap.L().Error("error while creating profile", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create profile")
	}

	return c.Status(fiber.StatusCreated).JSON(profile)
}

func (h *ProfileHandler) GetProfile(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		zap.L().Error("invalid user ID", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	profile, err := h.repo.GetProfile(c.Context(), int64(userID))
	if err != nil {
		zap.L().Error("error while getting profile", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve profile")
	}

	if profile == nil {
		zap.L().Error("profile not found", zap.Int("user_id", userID))
		return c.Status(fiber.StatusNotFound).SendString("Profile not found")
	}

	return c.Status(fiber.StatusOK).JSON(profile)
}

func (h *ProfileHandler) UpdateProfile(c *fiber.Ctx) error {
	req := new(request.UpdateProfileRequest)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("Validation failed: " + err.Error())
	}

	profile := &entity.Profile{
		Bio:               req.Bio,
		ProfilePictureURL: req.ProfilePicture,
		Experience:        req.Experience,
		Specialization:    req.Specialization,
		Photos:            req.Photos,
	}

	if err := h.repo.UpdateProfile(c.Context(), profile); err != nil {
		zap.L().Error("error while updating profile", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update profile")
	}

	return c.Status(fiber.StatusOK).JSON(profile)
}

func (h *ProfileHandler) UploadPhoto(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		zap.L().Error("invalid user ID", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	file, err := c.FormFile("photo")
	if err != nil {
		zap.L().Error("no file is received", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("No file is received")
	}

	// Save the file to a specific directory
	filePath := fmt.Sprintf("uploads/%d/%s", userID, file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		zap.L().Error("failed to upload file", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to upload file")
	}

	// Update the user's profile with the new photo URL
	if err := h.repo.AddPhoto(c.Context(), int64(userID), filePath); err != nil {
		zap.L().Error("failed to update profile with new photo", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update profile with new photo")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Photo uploaded successfully", "photo_url": filePath})
}
