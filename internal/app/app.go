package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/uxsnap/review_bot/internal/migrator"
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
	a.RunMigrations(ctx)
}

func (a *App) RunBotServer(ctx context.Context) {
	handlers := a.serviceProvider.Handlers(ctx)
	commands := make([]tele.Command, len(handlers))

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

	fmt.Println(whiteListIds)

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

func (a *App) RunMigrations(ctx context.Context) {
	migrator.Migrate(
		a.serviceProvider.dbClient,
	)
}
