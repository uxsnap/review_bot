package repositoryUsers

import (
	db "github.com/uxsnap/review_bot/internal/client/database"
	repositoryBase "github.com/uxsnap/review_bot/internal/repository"
)

type UsersRepository struct {
	*repositoryBase.BasePgRepository
}

func New(client db.DbClient) *UsersRepository {
	return &UsersRepository{
		repositoryBase.New(client),
	}
}
