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

func (cr *CategoriesRepository) Get(ctx context.Context, userID int64, name string, limit int, offset int) ([]entity.Category, error) {
	log.Printf("CategoriesRepository.Get, name: %v", name)

	var categories []entity.Category

	query := cr.DB().Where("user_id = ?", userID).Limit(limit).Offset(offset)

	if name != "" {
		query = query.Where("name LIKE ?", []string{name})
	}

	res := query.Find(&categories)

	if res.Error != nil {
		return []entity.Category{}, res.Error
	}

	return categories, nil
}

func (cr *CategoriesRepository) Count(ctx context.Context, userID int64) (int, error) {
	log.Printf("CategoriesRepository.Count, userID: %v", userID)

	var count int64
	err := cr.DB().Where("user_id = ?", userID).Count(&count).Error

	return int(count), err
}

func (cr *CategoriesRepository) Add(ctx context.Context, userID int64, name string, desc string) error {
	log.Printf("CategoriesRepository.Add, name: %v, desc: %v", name, desc)

	return cr.DB().Create(&entity.Category{
		Name:        name,
		Description: desc,
		UserID:      uint(userID),
	}).Error
}

func (cr *CategoriesRepository) Del(ctx context.Context, userID int64, name string) error {
	log.Printf("CategoriesRepository.Del, name: %v", name)

	return cr.DB().Debug().Where("name = ? and user_id = ?", name, userID).Delete(&entity.Category{}).Error
}
