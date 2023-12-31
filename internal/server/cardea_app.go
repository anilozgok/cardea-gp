package server

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/handlers"
)

type CardeaApp struct {
	register *handlers.RegisterHandler
	login    *handlers.LoginHandler
	getUsers *handlers.GetUsersHandler
}

func NewCardeaApp(repo database.Repository) *CardeaApp {
	return &CardeaApp{
		register: handlers.NewRegisterHandler(repo),
		login:    handlers.NewLoginHandler(repo),
		getUsers: handlers.NewGetUsersHandler(repo),
	}
}
