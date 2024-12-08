package app

import (
	"context"
	"log"
	"os"

	db "github.com/uxsnap/review_bot/internal/client/database"
	"github.com/uxsnap/review_bot/internal/client/database/sqlite"
)

type serviceProvider struct {
	dbClient db.DbClient
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
