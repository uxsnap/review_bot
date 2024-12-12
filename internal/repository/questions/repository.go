package repositoryQuestions

import (
	"context"
	"log"

	db "github.com/uxsnap/review_bot/internal/client/database"
	"github.com/uxsnap/review_bot/internal/entity"
	repositoryBase "github.com/uxsnap/review_bot/internal/repository"
)

type QuestionsRepository struct {
	*repositoryBase.BasePgRepository
}

func New(client db.DbClient) *QuestionsRepository {
	return &QuestionsRepository{
		repositoryBase.New(client),
	}
}

func (cr *QuestionsRepository) Add(ctx context.Context, categoryID int64, text string, answer string) error {
	log.Printf("QuestionsRepository.Add, name: %v, desc: %v", categoryID, text)

	return cr.DB().Create(&entity.Question{
		CategoryID: uint(categoryID),
		Text:       text,
		AnswerJSON: answer,
	}).Error
}
