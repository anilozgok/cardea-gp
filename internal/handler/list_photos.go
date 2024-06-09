package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type ListPhotosHandler struct {
	repo database.Repository
}

func NewListPhotosHandler(repo database.Repository) *ListPhotosHandler {
	return &ListPhotosHandler{repo: repo}
}

func (h *ListPhotosHandler) GetPhotosOfUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	photos, err := h.repo.GetImages(c.Context())
	if err != nil {
		zap.L().Error("error while getting images", zap.Uint("userId", userId), zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	usersPhotos := lo.Filter(photos, func(p entity.Photo, _ int) bool {
		return p.UserId == userId
	})

	photoURLs := lo.Map(usersPhotos, func(p entity.Photo, _ int) response.PhotoResponse {
		return response.PhotoResponse{
			PhotoId:   p.ID,
			PhotoURL:  fmt.Sprintf("%s/%s", c.BaseURL(), p.PhotoPath),
			CreatedAt: p.CreatedAt,
		}
	})

	return c.JSON(response.PhotosResponse{Photos: photoURLs})
}

func (h *ListPhotosHandler) GetPhotosOfStudents(c *fiber.Ctx) error {
	userIdParam := c.Params("userId")
	if userIdParam == "" {
		zap.L().Error("userId parameter is missing")
		return c.Status(fiber.StatusBadRequest).SendString("userId parameter is missing")
	}

	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		zap.L().Error("invalid userId parameter", zap.String("userIdParam", userIdParam), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).SendString("invalid userId parameter")
	}

	photos, err := h.repo.GetImages(c.Context())
	if err != nil {
		zap.L().Error("error while getting images", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	photosOfUser := lo.Filter(photos, func(p entity.Photo, _ int) bool {
		return p.UserId == uint(userId)
	})

	photoURLs := lo.Map(photosOfUser, func(p entity.Photo, _ int) response.PhotoResponse {
		return response.PhotoResponse{
			PhotoId:   p.ID,
			PhotoURL:  fmt.Sprintf("%s/%s", c.BaseURL(), p.PhotoPath),
			CreatedAt: p.CreatedAt,
		}
	})

	return c.JSON(response.PhotosResponse{Photos: photoURLs})
}

func (h *ListPhotosHandler) GetProfilePicture(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	images, err := h.repo.GetImages(c.Context())
	if err != nil {
		zap.L().Error("error while getting profile picture", zap.Uint("userId", userId), zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	usersProfilePicture := lo.Filter(images, func(p entity.Photo, _ int) bool {
		return p.UserId == userId && strings.HasPrefix(p.PhotoName, "pp_")
	})

	if len(usersProfilePicture) == 0 {
		return c.Status(fiber.StatusNotFound).SendString("profile picture not found")
	}

	profilePicture := usersProfilePicture[0]

	return c.JSON(response.PhotoResponse{
		PhotoId:   profilePicture.ID,
		PhotoURL:  fmt.Sprintf("%s/%s", c.BaseURL(), profilePicture.PhotoPath),
		CreatedAt: profilePicture.CreatedAt,
	})
}
