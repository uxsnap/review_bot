package app

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tele "gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

type App struct {
	Bot             *tele.Bot
	serviceProvider *serviceProvider
}

func New() (*App, error) {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)

	if err != nil {
		return nil, err
	}

	return &App{
		Bot:             bot,
		serviceProvider: newServiceProvider(),
	}, nil
}

func (a *App) Run(ctx context.Context) {
	a.RunBotServer(ctx)
}

func (a *App) RunBotServer(ctx context.Context) {
	handlers := a.serviceProvider.Handlers(ctx)
	commands := []tele.Command{}

	ids := strings.Split(os.Getenv("WHITELIST_IDS"), " ")

	whiteListIds := make([]int64, len(ids))

	for ind, id := range ids {
		conv, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Whitelist error", err)
			return
		}

		whiteListIds[ind] = int64(conv)
	}

	a.Bot.Use(middleware.Logger())
	a.Bot.Use(middleware.AutoRespond())
	a.Bot.Use(middleware.Whitelist(whiteListIds...))

	for endpoint, handler := range handlers {
		a.Bot.Handle(endpoint, handler)
		commands = append(commands, tele.Command{Text: endpoint, Description: endpoint})
	}

	a.Bot.SetCommands(commands)

	log.Println("\n === Bot has started working. === ")

	go a.Bot.Start()

	<-ctx.Done()

	log.Println("\n === Bot has stopped working. === ")
	a.Bot.Stop()
}
