package useCaseCategories

import (
	"context"

	"github.com/uxsnap/review_bot/internal/entity"
)

type CategoriesRepository interface {
	Get(ctx context.Context, name string) ([]entity.Category, error)
	Add(ctx context.Context, name string, desc string) error
	Del(ctx context.Context, name string) error
}
