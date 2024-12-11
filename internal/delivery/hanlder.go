package delivery

import (
	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	categoriesSubrouter "github.com/uxsnap/review_bot/internal/delivery/subrouters/categories"
	usersSubrouter "github.com/uxsnap/review_bot/internal/delivery/subrouters/users"
	"gopkg.in/telebot.v4"
)

func New(
	usersService subrouters.UsersService,
	categoriesService subrouters.CategoriesService,
) map[string]telebot.HandlerFunc {
	deps := subrouters.SubrouterDeps{
		UsersService:      usersService,
		CategoriesService: categoriesService,
	}

	handlers := map[string]map[string]telebot.HandlerFunc{
		"users":      usersSubrouter.New(deps),
		"categories": categoriesSubrouter.New(deps),
	}

	return prepareHandlers(handlers)
}

func prepareHandlers(handlers map[string]map[string]telebot.HandlerFunc) map[string]telebot.HandlerFunc {
	res := map[string]telebot.HandlerFunc{}

	for mainEndpoint, handlerMap := range handlers {
		for handlerEndpoint, handler := range handlerMap {
			curEndpoint := handlerEndpoint

			if curEndpoint != "" {
				curEndpoint = "_" + curEndpoint
			}

			res["/"+mainEndpoint+curEndpoint] = handler
		}
	}

	return res
}
