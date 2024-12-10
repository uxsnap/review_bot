package app

import (
	"context"
	"log"
	"os"

	db "github.com/uxsnap/review_bot/internal/client/database"
	"github.com/uxsnap/review_bot/internal/client/database/sqlite"
	"github.com/uxsnap/review_bot/internal/delivery"
	repositoryUsers "github.com/uxsnap/review_bot/internal/repository/users"
	ucUsers "github.com/uxsnap/review_bot/internal/usecase/users"
	"gopkg.in/telebot.v4"
)

type serviceProvider struct {
	dbClient db.DbClient
	handlers map[string]telebot.HandlerFunc

	usersRepository *repositoryUsers.UsersRepository
	ucUsers         *ucUsers.UseCaseUsers
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) SqliteClient(ctx context.Context) db.DbClient {
	if sp.dbClient == nil {
		client, err := sqlite.NewClient(ctx, os.Getenv("SQLITE_DB_NAME"))
		if err != nil {
			log.Fatalf("failed to connect to postgres: %v", err)
		}
		sp.dbClient = client
	}
	return sp.dbClient
}

func (sp *serviceProvider) UsersRepository(ctx context.Context) *repositoryUsers.UsersRepository {
	if sp.usersRepository == nil {
		sp.usersRepository = repositoryUsers.New(sp.SqliteClient(ctx))
	}
	return sp.usersRepository
}

func (sp *serviceProvider) UsersService(ctx context.Context) *ucUsers.UseCaseUsers {
	if sp.ucUsers == nil {
		sp.ucUsers = ucUsers.New(sp.UsersRepository(ctx))
	}
	return sp.ucUsers
}

func (sp *serviceProvider) Handlers(ctx context.Context) map[string]telebot.HandlerFunc {
	if len(sp.handlers) != 0 {
		return sp.handlers
	}

	sp.handlers = delivery.New(
		sp.UsersService(ctx),
	)

	// sp.handlers = map[string]telebot.HandlerFunc{
	// 	"/c": func(c telebot.Context) error {
	// 		return c.Send("Hello!")
	// 	},
	// }

	return sp.handlers
}
