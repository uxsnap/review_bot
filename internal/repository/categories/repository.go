package repositoryCategories

import (
	"context"
	"log"

	db "github.com/uxsnap/review_bot/internal/client/database"
	"github.com/uxsnap/review_bot/internal/entity"
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

func (cr *CategoriesRepository) Get(ctx context.Context, name string) ([]entity.Category, error) {
	log.Printf("CategoriesRepository.Get, name: %v", name)

	var categories []entity.Category

	query := cr.DB()

	if name != "" {
		query = query.Where("name LIKE ?", []string{name})
	}

	res := query.Find(&categories)

	if res.Error != nil {
		return []entity.Category{}, res.Error
	}

	return categories, nil
}

func (cr *CategoriesRepository) Add(ctx context.Context, name string, desc string) error {
	log.Printf("CategoriesRepository.Add, name: %v, desc: %v", name, desc)

	return cr.DB().Create(&entity.Category{
		Name:        name,
		Description: desc,
	}).Error
}

func (cr *CategoriesRepository) Del(ctx context.Context, name string) error {
	log.Printf("CategoriesRepository.Del, name: %v", name)

	return cr.DB().Debug().Where("name = ?", name).Delete(&entity.Category{}).Error
}
