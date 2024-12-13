package useCaseQuestions

import (
	"context"
	"log"
)

type UseCaseQuestions struct {
	questionsRepository QuestionsRepository
}

func New(
	questionsRepository QuestionsRepository,
) *UseCaseQuestions {
	return &UseCaseQuestions{
		questionsRepository: questionsRepository,
	}
}

func (uc *UseCaseQuestions) Add(ctx context.Context, categoryID int64, text string, answer string) error {
	log.Printf("UseCaseQuestions.Add, category: %v, text: %v, answer: %v", categoryID, text, answer)

	return uc.questionsRepository.Add(ctx, categoryID, text, answer)
}
