package handler

import (
	"fmt"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/model/response"
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
	userId := c.Locals("userId").(uint)

	req := new(request.CreateProfileRequest)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	maybeProfile, err := h.repo.GetProfileByUserId(c.Context(), userId)
	if err != nil {
		zap.L().Error("error while getting profile", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if maybeProfile != nil {
		zap.L().Error("profile already exists", zap.Uint("userId", userId))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	profile := &entity.Profile{
		UserId:            userId,
		Bio:               req.Bio,
		Height:            req.Height,
		Weight:            req.Weight,
		ProfilePictureURL: req.ProfilePicture,
		Experience:        req.Experience,
		Specialization:    req.Specialization,
	}

	if err := h.repo.CreateProfile(c.Context(), profile); err != nil {
		zap.L().Error("error while creating profile", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *ProfileHandler) UpdateProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	req := new(request.UpdateProfileRequest)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	profile, err := h.repo.GetProfileByUserId(c.Context(), userId)
	if err != nil {
		zap.L().Error("error while getting profile", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if profile == nil {
		zap.L().Error("profile not found", zap.Uint("userId", userId))
		return c.SendStatus(fiber.StatusNotFound)
	}

	profile.Bio = req.Bio
	profile.Height = req.Height
	profile.Weight = req.Weight
	profile.ProfilePictureURL = req.ProfilePicture
	profile.Experience = req.Experience
	profile.Specialization = req.Specialization

	if err := h.repo.UpdateProfile(c.Context(), profile); err != nil {
		zap.L().Error("error while updating profile", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update profile")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *ProfileHandler) GetProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	profile, err := h.repo.GetProfileByUserId(c.Context(), userId)
	if err != nil {
		zap.L().Error("error while getting profile", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if profile == nil {
		zap.L().Error("profile not found", zap.Uint("userId", userId))
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(response.ProfileResponse{
		UserId:         profile.UserId,
		Bio:            profile.Bio,
		Height:         profile.Height,
		Weight:         profile.Weight,
		ProfilePicture: profile.ProfilePictureURL,
		Experience:     profile.Experience,
		Specialization: profile.Specialization,
	})
}

func (h *ProfileHandler) UploadPhoto(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	file, err := c.FormFile("image")
	if err != nil {
		zap.L().Error("no file is received", zap.Error(err))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	filePath := fmt.Sprintf("uploads/%d/%s", userId, file.Filename)

	if err = c.SaveFile(file, filePath); err != nil {
		zap.L().Error("failed to upload file", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	photo := &entity.Image{
		UserId:    userId,
		ImageName: file.Filename,
		ImagePath: filePath,
	}

	if err = h.repo.AddPhoto(c.Context(), photo); err != nil {
		zap.L().Error("failed to update profile with new photo", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
