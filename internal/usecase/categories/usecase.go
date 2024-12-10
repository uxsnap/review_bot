package useCaseCategories

import (
	"context"

	"github.com/uxsnap/review_bot/internal/entity"
)

type UseCaseCategories struct {
	usersRepository CategoriesRepository
}

func New(
	usersRepository CategoriesRepository,
) *UseCaseCategories {
	return &UseCaseCategories{
		usersRepository: usersRepository,
	}
}

func (uc *UseCaseCategories) Get(ctx context.Context, name string) ([]entity.Category, error) {
	return []entity.Category{}, nil
}
