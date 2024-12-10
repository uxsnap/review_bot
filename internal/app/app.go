package app

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/uxsnap/review_bot/internal/migrator"
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
	a.RunMigrations(ctx)
}

func (a *App) RunBotServer(ctx context.Context) {
	handlers := a.serviceProvider.Handlers(ctx)

	a.Bot.SetCommands([]tele.Command{
		{Text: "/users", Description: "get users"},
	})

	// fmt.Println(handlers)

	for endpoint, handler := range handlers {
		a.Bot.Handle(endpoint, handler)
	}

	log.Println("\n === Bot has started working. === ")

	a.Bot.Start()

	go func() {
		<-ctx.Done()

		a.Bot.Stop()
		log.Println("\n === Bot has stopped working. === ")
	}()
}

func (a *App) RunMigrations(ctx context.Context) {
	migrator.Migrate(
		a.serviceProvider.dbClient,
	)
}
