package questionsSubrouter

import (
	"context"
	"log"
	"strconv"

	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	"gopkg.in/telebot.v4"
)

func (cs *QuestionsSubrouter) addQuestion(tctx telebot.Context) error {
	log.Println("called: addQuestion")

	ctx := context.Background()

	categories, err := cs.CategoriesService.Get(ctx, tctx.Update().Message.Sender.ID, "", subrouters.LIMIT_COUNT, 0)

	if err != nil {
		log.Printf("error: addQuestion, %v", err)
		return tctx.Send("Не удалось получить категории :С")
	}

	if len(categories) == 0 {
		log.Printf("error: addQuestion, %v", err)
		return tctx.Send("Нет категорий :С")
	}

	categoryRows := []telebot.Row{}
	selector := &telebot.ReplyMarkup{}

	for _, c := range categories {
		conv := strconv.Itoa(int(c.ID))

		categoryRows = append(categoryRows, selector.Row(selector.Data(c.Name, "addQuestion", conv)))
	}

	categoryRows = append(categoryRows, selector.Row(
		selector.Data("➡", "addQuestion", "button_next", "1"),
	))

	selector.Inline(categoryRows...)

	return tctx.Send("Выберите категорию вопроса: ", selector)
}
