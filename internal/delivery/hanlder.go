package delivery

import (
	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	categoriesSubrouter "github.com/uxsnap/review_bot/internal/delivery/subrouters/categories"
	messageHandlersSubrouter "github.com/uxsnap/review_bot/internal/delivery/subrouters/messageHandlers"
	questionsSubrouter "github.com/uxsnap/review_bot/internal/delivery/subrouters/questions"
	usersSubrouter "github.com/uxsnap/review_bot/internal/delivery/subrouters/users"
	"gopkg.in/telebot.v4"
)

func New(
	kvClient subrouters.KvClient,
	usersService subrouters.UsersService,
	categoriesService subrouters.CategoriesService,
	questionsService subrouters.QuestionsService,
) map[interface{}]telebot.HandlerFunc {
	deps := subrouters.SubrouterDeps{
		KvClient:          kvClient,
		UsersService:      usersService,
		CategoriesService: categoriesService,
		QuestionsService:  questionsService,
	}

	mainHandlers := prepareHandlers(map[interface{}]map[string]telebot.HandlerFunc{
		"users":      usersSubrouter.New(deps),
		"categories": categoriesSubrouter.New(deps),
		"questions":  questionsSubrouter.New(deps),
	})

	messageHandlers := map[interface{}]telebot.HandlerFunc{
		telebot.OnCallback: messageHandlersSubrouter.Handle(deps),
		telebot.OnText:     messageHandlersSubrouter.Handle(deps),
	}

	handlers := map[interface{}]telebot.HandlerFunc{}

	for k, v := range mainHandlers {
		handlers[k] = v
	}

	for k, v := range messageHandlers {
		handlers[k] = v
	}

	return handlers
}

func prepareHandlers(handlers map[interface{}]map[string]telebot.HandlerFunc) map[string]telebot.HandlerFunc {
	res := map[string]telebot.HandlerFunc{}

	for mainEndpoint, handlerMap := range handlers {
		for handlerEndpoint, handler := range handlerMap {
			curEndpoint := handlerEndpoint

			if curEndpoint != "" {
				curEndpoint = "_" + curEndpoint
			}

			res["/"+mainEndpoint.(string)+curEndpoint] = handler
		}
	}

	return res
}
