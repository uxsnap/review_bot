package delivery

import (
	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	usersSubrouter "github.com/uxsnap/review_bot/internal/delivery/subrouters/users"
	"gopkg.in/telebot.v4"
)

func New(
	usersService subrouters.UsersService,
) map[string]telebot.HandlerFunc {
	deps := subrouters.SubrouterDeps{
		UsersService: usersService,
	}

	handlers := map[string]map[string]telebot.HandlerFunc{
		"/users": usersSubrouter.New(deps),
	}

	return prepareHandlers(handlers)
}

func prepareHandlers(handlers map[string]map[string]telebot.HandlerFunc) map[string]telebot.HandlerFunc {
	res := map[string]telebot.HandlerFunc{}

	for mainEndpoint, handlerMap := range handlers {
		for curEndpoint, handler := range handlerMap {
			if curEndpoint == "/" {
				curEndpoint = ""
			}

			res[mainEndpoint+curEndpoint] = handler
		}
	}

	return res
}