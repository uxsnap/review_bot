package usersSubrouter

import (
	"fmt"

	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	"gopkg.in/telebot.v4"
)

type UsersSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) map[string]telebot.HandlerFunc {
	// us := UsersSubrouter{deps}

	return map[string]telebot.HandlerFunc{
		"": func(c telebot.Context) error {
			err := c.Send("USers!")

			if err != nil {
				fmt.Println(err)
			}

			return nil
		},
	}
}
