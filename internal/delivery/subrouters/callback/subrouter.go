package callbackSubrouter

import (
	"strings"

	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	"gopkg.in/telebot.v4"
)

type CallbackSubrouter struct {
	subrouters.SubrouterDeps
}

func Handle(deps subrouters.SubrouterDeps) telebot.HandlerFunc {
	cs := CallbackSubrouter{deps}

	callbackMapping := map[string]telebot.HandlerFunc{
		"addQuestion": cs.addQuestion,
	}

	return func(ctx telebot.Context) error {
		data := strings.Split(ctx.Data(), "|")
		queryName := data[0]

		return callbackMapping[queryName](ctx)
	}
}
