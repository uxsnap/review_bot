package useCaseQuestions

import (
	"context"
)

type QuestionsRepository interface {
	Add(ctx context.Context, categoryID int64, text string, answer string) error
}
