package subrouters

import (
	"context"

	"github.com/uxsnap/review_bot/internal/entity"
)

type UsersService interface {
}

type CategoriesService interface {
	Get(ctx context.Context, name string) ([]entity.Category, error)
}
