package app

import (
	"context"
	"log"
	"os"

	db "github.com/uxsnap/review_bot/internal/client/database"
	kv "github.com/uxsnap/review_bot/internal/client/database/keyValue"
	"github.com/uxsnap/review_bot/internal/client/database/sqlite"
	"github.com/uxsnap/review_bot/internal/delivery"
	"github.com/uxsnap/review_bot/internal/migrator"
	repositoryCategories "github.com/uxsnap/review_bot/internal/repository/categories"
	repositoryQuestions "github.com/uxsnap/review_bot/internal/repository/questions"
	repositoryUsers "github.com/uxsnap/review_bot/internal/repository/users"
	ucCategories "github.com/uxsnap/review_bot/internal/usecase/categories"
	ucQuestions "github.com/uxsnap/review_bot/internal/usecase/questions"
	ucUsers "github.com/uxsnap/review_bot/internal/usecase/users"
	"gopkg.in/telebot.v4"
)

type serviceProvider struct {
	dbClient db.DbClient
	kvClient db.KvClient
	handlers map[interface{}]telebot.HandlerFunc

	usersRepository      *repositoryUsers.UsersRepository
	categoriesRepository *repositoryCategories.CategoriesRepository
	questionsRepository  *repositoryQuestions.QuestionsRepository

	ucUsers      *ucUsers.UseCaseUsers
	ucCategories *ucCategories.UseCaseCategories
	ucQuestions  *ucQuestions.UseCaseQuestions
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) SqliteClient(ctx context.Context) db.DbClient {
	if sp.dbClient == nil {
		client, err := sqlite.NewClient(ctx, os.Getenv("SQLITE_DB_NAME"))
		if err != nil {
			log.Fatalf("failed to connect to sqlite: %v", err)
		}

		migrator.Migrate(
			client,
		)

		sp.dbClient = client
	}
	return sp.dbClient
}

func (sp *serviceProvider) MapKvClient(ctx context.Context) db.KvClient {
	if sp.kvClient == nil {
		client, err := kv.NewKvClient(ctx)
		if err != nil {
			log.Fatalf("failed to connect to kv client: %v", err)
		}

		sp.kvClient = client
	}
	return sp.kvClient
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

func (sp *serviceProvider) QuestionsRepository(ctx context.Context) *repositoryQuestions.QuestionsRepository {
	if sp.questionsRepository == nil {
		sp.questionsRepository = repositoryQuestions.New(sp.SqliteClient(ctx))
	}
	return sp.questionsRepository
}

func (sp *serviceProvider) QuestionsService(ctx context.Context) *ucQuestions.UseCaseQuestions {
	if sp.ucQuestions == nil {
		sp.ucQuestions = ucQuestions.New(sp.QuestionsRepository(ctx))
	}
	return sp.ucQuestions
}

func (sp *serviceProvider) Handlers(ctx context.Context) map[interface{}]telebot.HandlerFunc {
	if len(sp.handlers) != 0 {
		return sp.handlers
	}

	sp.handlers = delivery.New(
		sp.MapKvClient(ctx),
		sp.UsersService(ctx),
		sp.CategoriesService(ctx),
		sp.QuestionsService(ctx),
	)

	return sp.handlers
}
