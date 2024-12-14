package useCaseCategories

import (
	"context"

	"github.com/uxsnap/review_bot/internal/entity"
)

type CategoriesRepository interface {
	Get(ctx context.Context, userID int64, name string, limit int, offset int) ([]entity.Category, error)
	Count(ctx context.Context, userID int64) (int, error)
	Add(ctx context.Context, userID int64, name string, desc string) error
	Del(ctx context.Context, userID int64, name string) error
}
