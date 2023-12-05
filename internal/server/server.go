package server

import (
	"github.com/anilozgok/cardea-gp/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppServer struct {
	app *fiber.App
}

func NewAppServer(db *gorm.DB) *AppServer {
	app := fiber.New()

	//app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	//TODO::define liveness and readiness probes to this endpoint while deploying to k8s
	app.Get("/health", healthCheck)

	r := app.Group("/api/v1")

	//USER ENDPOINTS
	userGroup := r.Group("/users")
	userGroup.Post("/", handlers.CreateNewUserHandler(db))

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
