package useCaseCategories

import (
	"context"
	"log"

	"github.com/uxsnap/review_bot/internal/entity"
)

type UseCaseCategories struct {
	categoriesRepository CategoriesRepository
}

func New(
	categoriesRepository CategoriesRepository,
) *UseCaseCategories {
	return &UseCaseCategories{
		categoriesRepository: categoriesRepository,
	}
}

func (uc *UseCaseCategories) Get(ctx context.Context, name string) ([]entity.Category, error) {
	log.Printf("UseCaseCategories.Get, name: %v", name)

	return uc.categoriesRepository.Get(ctx, name)
}

func (uc *UseCaseCategories) Add(ctx context.Context, name string, desc string) error {
	log.Printf("UseCaseCategories.Add, name: %v, desc: %v", name, desc)

	return uc.categoriesRepository.Add(ctx, name, desc)
}

func (uc *UseCaseCategories) Del(ctx context.Context, name string) error {
	log.Printf("UseCaseCategories.Del, name: %v", name)

	return uc.categoriesRepository.Del(ctx, name)
}
