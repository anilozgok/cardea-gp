package handler

import (
	"fmt"

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
	coachId := c.Locals("userId").(uint)

	usersOfCoach, err := h.repo.GetStudentsOfCoach(c.Context(), coachId)
	if err != nil {
		zap.L().Error("error while getting students", zap.Uint("coachId", coachId), zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	photos, err := h.repo.GetImages(c.Context())
	if err != nil {
		zap.L().Error("error while getting images", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userIds := lo.Map(usersOfCoach, func(u entity.User, _ int) uint {
		return u.ID
	})

	photosOfStudents := lo.Filter(photos, func(p entity.Photo, _ int) bool {
		return lo.Contains(userIds, p.UserId)
	})

	photoURLs := lo.Map(photosOfStudents, func(p entity.Photo, _ int) response.PhotoResponse {
		return response.PhotoResponse{
			PhotoId:   p.ID,
			PhotoURL:  fmt.Sprintf("%s/%s", c.BaseURL(), p.PhotoPath),
			CreatedAt: p.CreatedAt,
		}
	})

	return c.JSON(response.PhotosResponse{Photos: photoURLs})
}
