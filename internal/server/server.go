package server

import (
	"github.com/anilozgok/cardea-gp/internal/config"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppServer struct {
	app    *fiber.App
	config *config.Config
}

func NewAppServer(db *gorm.DB, isLocalMode bool) *AppServer {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	if isLocalMode {
		app.Use(logger.New())
	}

	app.Get("/health", healthCheck)

	r := app.Group("/api/v1")
	repo := database.NewRepository(db)

	cardeaApp := NewCardeaApp(repo)

	router := NewRouter(cardeaApp, r)
	router.InitializeRoute()

	return &AppServer{
		app: app,
	}
}

func (s *AppServer) Start() {
	go func() {
		err := s.app.Listen(":8080")
		if err != nil {
			zap.L().Fatal("error while starting server", zap.Error(err))
		}
	}()
}

func (s *AppServer) Shutdown() error {
	return s.app.Shutdown()
}

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
