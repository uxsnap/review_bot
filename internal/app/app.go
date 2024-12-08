package app

import (
	"context"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
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
	go func() {
		a.Bot.Start()

		log.Println("\n === Bot has started working. === ")
	}()

	go func() {
		<-ctx.Done()

		log.Println("\n === Bot is shutting down. === ")

		a.Bot.Stop()

		log.Println("\n === Bot has shut down. === ")
	}()
}
