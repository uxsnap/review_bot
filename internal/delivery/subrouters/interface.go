package subrouters

import (
	"context"

	"github.com/uxsnap/review_bot/internal/entity"
)

type UsersService interface {
}

type CategoriesService interface {
	Get(ctx context.Context, userID int64, name string, limit int, offset int) ([]entity.Category, error)
	Add(ctx context.Context, userID int64, name string, desc string) error
	Del(ctx context.Context, userID int64, name string) error
}

type QuestionsService interface {
	Add(ctx context.Context, categoryID int64, text string, answer string) error
}
