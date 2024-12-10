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
	a.serviceProvider.SqliteClient(ctx)

	a.RunBotServer(ctx)
	a.RunMigrations(ctx)
}

func (a *App) RunBotServer(ctx context.Context) {
	handlers := a.serviceProvider.Handlers(ctx)

	for endpoint, handler := range handlers {
		a.Bot.Handle(endpoint, handler)
	}

	log.Println("\n === Bot has started working. === ")

	a.Bot.Start()
}

func (a *App) RunMigrations(ctx context.Context) {
	migrator.Migrate(
		a.serviceProvider.dbClient,
	)
}
