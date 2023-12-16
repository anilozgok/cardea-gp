package server

import (
	"github.com/anilozgok/cardea-gp/internal/config"
	"github.com/anilozgok/cardea-gp/internal/handlers"
	"github.com/anilozgok/cardea-gp/internal/middleware"
	"github.com/anilozgok/cardea-gp/internal/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppServer struct {
	app    *fiber.App
	config *config.Config
}

func NewAppServer(db *gorm.DB, config *config.Config) *AppServer {
	app := fiber.New()
	app.Get("/health", healthCheck)
	r := app.Group("/api/v1")

	repo := repository.NewRepository(db)

	register := handlers.NewRegisterHandler(repo)
	r.Post("/register", register.Handle)

	login := handlers.NewLoginHandler(repo)
	r.Post("/login", login.Handle)

	logout := handlers.NewLogoutHandler(repo)
	r.Post("/logout", middleware.AuthMiddleware, logout.Handle)

	userGroup := r.Group("/users")
	userGroup.Use(middleware.AuthMiddleware)

	return &AppServer{
		app: app,
	}
}

func (s *AppServer) Start() {
	go s.app.Listen(":8080")
}

func (s *AppServer) Shutdown() error {
	return s.app.Shutdown()
}

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
