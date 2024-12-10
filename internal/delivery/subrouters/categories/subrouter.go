package categoriesSubrouter

import (
	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	"gopkg.in/telebot.v4"
)

type CategoriesSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) map[string]telebot.HandlerFunc {
	router := CategoriesSubrouter{deps}

	return map[string]telebot.HandlerFunc{
		"/": router.getAllCategories,
	}
}
