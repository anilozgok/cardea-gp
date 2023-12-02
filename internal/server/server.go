package server

import (
	"github.com/anilozgok/cardea-gp/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

type AppServer struct {
	app *fiber.App
}

func NewAppServer() *AppServer {
	app := fiber.New()

	//app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	//TODO::define liveness and readiness probes to this endpoint while deploying to k8s
	app.Get("/health", healthCheck)

	r := app.Group("/api/v1")
	r.Get("/users", handlers.NewGetUsersHandler())

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
