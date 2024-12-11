package sqlite

import (
	"context"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbClient struct {
	dbc *gorm.DB
}

func NewClient(ctx context.Context, dbName string) (*DbClient, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer

		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,
		},
	)})

	if err != nil {
		panic("failed to connect database")
	}

	return &DbClient{
		dbc: db,
	}, nil
}

func (c *DbClient) DB() *gorm.DB {
	return c.dbc
}

func (c *DbClient) Close(ctx context.Context) error {
	if c.dbc != nil {
		dbInst, err := c.dbc.DB()

		if err != nil {
			panic("failed to connect to database")
		}

		dbInst.Close()
	}

	return nil
}
