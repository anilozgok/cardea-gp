package server

import (
	"github.com/anilozgok/cardea-gp/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	cardeaApp *CardeaApp
	router    fiber.Router
}

func NewRouter(cardeaApp *CardeaApp, router fiber.Router) *Router {
	return &Router{
		cardeaApp: cardeaApp,
		router:    router,
	}
}

func (r *Router) InitializeRoute() {
	auth := r.router.Group("/auth")
	auth.Post("/register", r.cardeaApp.register.Handle)
	auth.Post("/login", r.cardeaApp.login.Handle)

	user := r.router.Group("/user")
	user.Get("/", middleware.AuthMiddleware, middleware.RoleAdmin, r.cardeaApp.getUsers.Handle)

}
