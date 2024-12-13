package db

import (
	"context"

	"gorm.io/gorm"
)

type DbClient interface {
	DB() *gorm.DB
	Close(ctx context.Context) error
}

type KvClient interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
	Delete(key string)
}
