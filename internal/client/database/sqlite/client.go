package sqlite

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbClient struct {
	dbc *gorm.DB
}

func NewClient(ctx context.Context, dbName string) (*DbClient, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

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
