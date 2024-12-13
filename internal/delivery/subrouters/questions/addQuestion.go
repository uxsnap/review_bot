package questionsSubrouter

import (
	"context"
	"log"
	"strconv"

	"gopkg.in/telebot.v4"
)

func (cs *QuestionsSubrouter) addQuestion(tctx telebot.Context) error {
	log.Println("called: addQuestion")

	ctx := context.Background()

	categories, err := cs.CategoriesService.Get(ctx, tctx.Update().Message.Sender.ID, "", 1, 0)

	if err != nil {
		log.Printf("error: addQuestion, %v", err)
		return tctx.Send("Не удалось получить категории :С")
	}

	categoryRows := []telebot.Row{}
	selector := &telebot.ReplyMarkup{}

	for _, c := range categories {
		conv := strconv.Itoa(int(c.ID))

		categoryRows = append(categoryRows, selector.Row(selector.Data(c.Name, "addQuestion", conv)))
	}

	categoryRows = append(categoryRows, selector.Row(
		selector.Data("⬅", "addQuestion", "button_prev", "1"),
		selector.Data("➡", "addQuestion", "button_next", "2"),
	))

	selector.Inline(categoryRows...)

	return tctx.Send("Выберите категорию вопроса: ", selector)
}
