package callbackSubrouter

import (
	"context"
	"log"
	"strconv"
	"strings"

	"gopkg.in/telebot.v4"
)

func (cs *CallbackSubrouter) getQuestionCategories(tctx telebot.Context) error {
	ctx := context.Background()

	categories, err := cs.CategoriesService.Get(ctx, tctx.Update().Message.Sender.ID, "", 1, 1)

	if err != nil {
		log.Printf("error: addQuestion, %v", err)
		return tctx.Send("Не удалось получить категории :С")
	}

	categoryRows := []telebot.Row{}
	selector := &telebot.ReplyMarkup{}

	for _, c := range categories {
		conv := strconv.Itoa(int(c.ID))

		categoryRows = append(categoryRows, selector.Row(selector.Data(c.Name, conv, "addQuestion")))
	}

	categoryRows = append(categoryRows, selector.Row(
		selector.Data("⬅", "addQuestion", "prev", "1"),
		selector.Data("➡", "addQuestion", "next", "3"),
	))

	selector.Inline(categoryRows...)

	return tctx.Edit("Выберите категорию вопроса: ", selector)
}

func (cs *CallbackSubrouter) addQuestion(tctx telebot.Context) error {
	log.Println("called: addQuestion")

	data := strings.Split(tctx.Data(), "|")[1:] // remove addQuestion

	if data[0] == "button_next" || data[0] == "button_prev" {
		return cs.getQuestionCategories(tctx)
	}

	return tctx.Send("Ну выбрал ты блин")
}
