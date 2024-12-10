package repositoryCategories

import (
	db "github.com/uxsnap/review_bot/internal/client/database"
	repositoryBase "github.com/uxsnap/review_bot/internal/repository"
)

type CategoriesRepository struct {
	*repositoryBase.BasePgRepository
}

func New(client db.DbClient) *CategoriesRepository {
	return &CategoriesRepository{
		repositoryBase.New(client),
	}
}
