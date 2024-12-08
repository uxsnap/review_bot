package db

import (
	"context"

	"gorm.io/gorm"
)

type DbClient interface {
	DB() *gorm.DB
	Close(ctx context.Context) error
}
