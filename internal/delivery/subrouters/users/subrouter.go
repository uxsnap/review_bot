package usersSubrouter

import (
	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	"gopkg.in/telebot.v4"
)

type UsersSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) map[string]telebot.HandlerFunc {
	// us := UsersSubrouter{deps}

	return map[string]telebot.HandlerFunc{
		"/": func(ctx telebot.Context) error {
			return ctx.Send("Hello!")
		},
	}
}
