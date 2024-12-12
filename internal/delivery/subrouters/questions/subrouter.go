package questionsSubrouter

import (
	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	"gopkg.in/telebot.v4"
)

type QuestionsSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) map[string]telebot.HandlerFunc {
	router := QuestionsSubrouter{deps}

	return map[string]telebot.HandlerFunc{
		"add": router.addQuestion,
	}
}
