package callbackSubrouter

import (
	"fmt"
	"strings"

	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	"gopkg.in/telebot.v4"
)

type CallbackSubrouter struct {
	subrouters.SubrouterDeps
}

func Handle(deps subrouters.SubrouterDeps) telebot.HandlerFunc {
	router := CallbackSubrouter{deps}

	callbackMapping := map[string]telebot.HandlerFunc{
		"addQuestion": router.addQuestion,
	}

	return func(ctx telebot.Context) error {
		data := ctx.Args()
		queryName := strings.TrimSpace(data[0])

		fmt.Println(queryName)

		return callbackMapping[queryName](ctx)
	}
}
