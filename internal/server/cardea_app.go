package server

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/handler"
)

type CardeaApp struct {
	register *handler.RegisterHandler
	login    *handler.LoginHandler
	getUsers *handler.GetUsersHandler
	logout   *handler.LogoutHandler
	me       *handler.MeHandler
}

func NewCardeaApp(repo database.Repository) *CardeaApp {
	return &CardeaApp{
		register: handler.NewRegisterHandler(repo),
		login:    handler.NewLoginHandler(repo),
		getUsers: handler.NewGetUsersHandler(repo),
		logout:   handler.NewLogoutHandler(),
		me:       handler.NewMeHandler(),
	}
}
