package app

import (
	"context"
	"log"
	"os"

	db "github.com/uxsnap/review_bot/internal/client/database"
	"github.com/uxsnap/review_bot/internal/client/database/sqlite"
	"github.com/uxsnap/review_bot/internal/delivery"
	repositoryCategories "github.com/uxsnap/review_bot/internal/repository/categories"
	repositoryUsers "github.com/uxsnap/review_bot/internal/repository/users"
	ucCategories "github.com/uxsnap/review_bot/internal/usecase/categories"
	ucUsers "github.com/uxsnap/review_bot/internal/usecase/users"
	"gopkg.in/telebot.v4"
)

type serviceProvider struct {
	dbClient db.DbClient
	handlers map[string]telebot.HandlerFunc

	usersRepository      *repositoryUsers.UsersRepository
	categoriesRepository *repositoryCategories.CategoriesRepository

	ucUsers      *ucUsers.UseCaseUsers
	ucCategories *ucCategories.UseCaseCategories
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

func (sp *serviceProvider) CategoriesRepository(ctx context.Context) *repositoryCategories.CategoriesRepository {
	if sp.categoriesRepository == nil {
		sp.categoriesRepository = repositoryCategories.New(sp.SqliteClient(ctx))
	}
	return sp.categoriesRepository
}

func (sp *serviceProvider) CategoriesService(ctx context.Context) *ucCategories.UseCaseCategories {
	if sp.ucCategories == nil {
		sp.ucCategories = ucCategories.New(sp.CategoriesRepository(ctx))
	}
	return sp.ucCategories
}

func (sp *serviceProvider) Handlers(ctx context.Context) map[string]telebot.HandlerFunc {
	if len(sp.handlers) != 0 {
		return sp.handlers
	}

	sp.handlers = delivery.New(
		sp.UsersService(ctx),
		sp.CategoriesService(ctx),
	)

	return sp.handlers
}
