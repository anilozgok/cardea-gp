package server

import (
	"github.com/anilozgok/cardea-gp/internal/config"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppServer struct {
	app    *fiber.App
	config *config.Config
}

type CardeaApp struct {
	register *handlers.RegisterHandler
	login    *handlers.LoginHandler
	logout   *handlers.LogoutHandler
	getUsers *handlers.GetUsersHandler
}

func NewAppServer(db *gorm.DB) *AppServer {
	app := fiber.New()
	app.Get("/health", healthCheck)

	r := app.Group("/api/v1")
	repo := database.NewRepository(db)

	cardeaApp := &CardeaApp{
		register: handlers.NewRegisterHandler(repo),
		login:    handlers.NewLoginHandler(repo),
		logout:   handlers.NewLogoutHandler(repo),
		getUsers: handlers.NewGetUsersHandler(repo),
	}

	router := NewRouter(cardeaApp, r)
	router.InitializeRoute()

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
