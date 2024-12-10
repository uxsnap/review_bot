package repositoryBase

import (
	db "github.com/uxsnap/review_bot/internal/client/database"
	"gorm.io/gorm"
)

type BasePgRepository struct {
	dbc db.DbClient
}

func New(client db.DbClient) *BasePgRepository {
	return &BasePgRepository{
		dbc: client,
	}
}

func (r *BasePgRepository) DB() *gorm.DB {
	return r.dbc.DB()
}
