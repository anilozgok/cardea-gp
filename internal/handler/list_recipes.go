package handler

import (
	"encoding/json"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"io/ioutil"
	"path/filepath"
)

type ListRecipesHandler struct{}

func NewListRecipesHandler() *ListRecipesHandler {
	return &ListRecipesHandler{}
}

func (h *ListRecipesHandler) Handle(ctx *fiber.Ctx) error {
	filePath := filepath.Join("recipes_raw_nosource_epi.json")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		zap.L().Error("Failed to read recipes file", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to read recipes file")
	}

	var recipes map[string]entity.Recipe
	err = json.Unmarshal(file, &recipes)
	if err != nil {
		zap.L().Error("Failed to unmarshal recipes", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to unmarshal recipes")
	}

	recipesList := make([]entity.Recipe, 0, len(recipes))
	for id, recipe := range recipes {
		recipe.ID = id
		recipesList = append(recipesList, recipe)
	}

	return ctx.JSON(recipesList)
}
