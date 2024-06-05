package handler

import (
	"fmt"
	"os"

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
		Phone:             req.Specialization,
		Country:           req.Country,
		StateProvince:     req.StateProvince,
		Address:           req.Address,
		City:              req.Address,
		ZipCode:           req.ZipCode,
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
	profile.Address = req.Address
	profile.City = req.City
	profile.Country = req.Country
	profile.Phone = req.Phone
	profile.ZipCode = req.ZipCode
	profile.StateProvince = req.StateProvince

	if err := h.repo.UpdateProfile(c.Context(), profile); err != nil {
		zap.L().Error("error while updating profile", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update profile")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *ProfileHandler) GetProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	maybeUser, err := h.repo.GetUserById(c.Context(), userId)
	if err != nil {
		zap.L().Error("error while getting user", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if maybeUser == nil {
		zap.L().Error("user not found", zap.Uint("userId", userId))
		return c.SendStatus(fiber.StatusNotFound)
	}

	maybeProfile, err := h.repo.GetProfileByUserId(c.Context(), userId)
	if err != nil {
		zap.L().Error("error while getting profile", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if maybeProfile == nil {
		zap.L().Error("profile not found", zap.Uint("userId", userId))
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(response.ProfileResponse{
		FirstName:      maybeUser.FirstName,
		LastName:       maybeUser.LastName,
		Email:          maybeUser.Email,
		Gender:         maybeUser.Gender,
		DateOfBirth:    maybeUser.DateOfBirth,
		Bio:            maybeProfile.Bio,
		Height:         maybeProfile.Height,
		Weight:         maybeProfile.Weight,
		ProfilePicture: maybeProfile.ProfilePictureURL,
		Experience:     maybeProfile.Experience,
		Specialization: maybeProfile.Specialization,
		Phone:          maybeProfile.Phone,
		Country:        maybeProfile.Country,
		StateProvince:  maybeProfile.StateProvince,
		Address:        maybeProfile.Address,
		City:           maybeProfile.City,
		ZipCode:        maybeProfile.ZipCode,
	})
}

func (h *ProfileHandler) UploadPhoto(c *fiber.Ctx) error {
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err := os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	userId := c.Locals("userId").(uint)

	file, err := c.FormFile("image")
	if err != nil {
		zap.L().Error("no file is received", zap.Error(err))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fileName := fmt.Sprintf("%d_%s", userId, file.Filename)
	filePath := fmt.Sprintf("./uploads/%s", fileName)

	if err = c.SaveFile(file, filePath); err != nil {
		zap.L().Error("failed to upload file", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	photo := &entity.Photo{
		UserId:    userId,
		PhotoName: fileName,
		PhotoPath: filePath,
	}

	if err = h.repo.AddPhoto(c.Context(), photo); err != nil {
		zap.L().Error("failed to update profile with new photo", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
