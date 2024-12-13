package messageHandlersSubrouter

import (
	"strconv"
	"strings"

	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	"gopkg.in/telebot.v4"
)

type MessageHandlersSubrouter struct {
	subrouters.SubrouterDeps
}

func Handle(deps subrouters.SubrouterDeps) telebot.HandlerFunc {
	router := MessageHandlersSubrouter{deps}

	messageHandlersMapping := map[string]telebot.HandlerFunc{
		"addQuestion":        router.addQuestion,
		"handleTextQuestion": router.handleTextQuestion,
		"handleTextAnswer":   router.handleTextAnswer,
	}

	return func(tctx telebot.Context) error {
		data := tctx.Args()

		var queryName string

		if len(data) == 0 {
			sender := tctx.Sender()

			lastUserAction, ok := router.KvClient.Get(strconv.Itoa(int(sender.ID)))
			if !ok {
				return tctx.Send("Не удалось получить пользователя :С")
			}

			userQI := lastUserAction.(*UserQuestionInfo)

			tctx.Set(userQI.actionType, userQI.data)

			queryName = strings.TrimSpace(userQI.actionType)

		} else {
			queryName = strings.TrimSpace(data[0])

		}

		return messageHandlersMapping[queryName](tctx)
	}
}
