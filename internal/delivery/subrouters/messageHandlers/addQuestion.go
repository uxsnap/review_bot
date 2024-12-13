package messageHandlersSubrouter

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"gopkg.in/telebot.v4"
)

type UserQuestionInfo struct {
	actionType string
	data       []string
}

func (cs *MessageHandlersSubrouter) addQuestion(tctx telebot.Context) error {
	log.Println("called: addQuestion callback")

	data := tctx.Args()[1:] // remove addQuestion
	query := data[0]

	if query == "button_next" || query == "button_prev" {
		return cs.getQuestionCategories(tctx)
	}

	if query == "questionType" {
		return cs.addTextQuestion(tctx)
	}

	selector := &telebot.ReplyMarkup{}

	selector.Inline(
		selector.Row(
			selector.Data("Вопрос", "addQuestion", append([]string{"questionType"}, data...)...),
		),
	)

	return tctx.Edit("Выбери тип вопроса:", selector)
}

func (cs *MessageHandlersSubrouter) getQuestionCategories(tctx telebot.Context) error {
	ctx := context.Background()

	categories, err := cs.CategoriesService.Get(ctx, tctx.Callback().Sender.ID, "", 1, 1)

	if err != nil {
		log.Printf("error: addQuestion addQuestion callback, %v", err)
		return tctx.Send("Не удалось получить категории :С")
	}

	categoryRows := []telebot.Row{}
	selector := &telebot.ReplyMarkup{}

	for _, c := range categories {
		conv := strconv.Itoa(int(c.ID))

		categoryRows = append(categoryRows, selector.Row(selector.Data(c.Name, "addQuestion", conv)))
	}

	categoryRows = append(categoryRows, selector.Row(
		selector.Data("⬅", "addQuestion", "prev", "1"),
		selector.Data("➡", "addQuestion", "next", "3"),
	))

	selector.Inline(categoryRows...)

	return tctx.Edit("Выберите категорию вопроса: ", selector)
}

func (cs *MessageHandlersSubrouter) addTextQuestion(tctx telebot.Context) error {
	log.Println("called: addTextQuestion")

	sender := tctx.Sender()

	cs.KvClient.Set(strconv.Itoa(int(sender.ID)), &UserQuestionInfo{
		actionType: "handleTextQuestion",
		data:       tctx.Args(),
	})

	return tctx.Edit("Введите вопрос: ")
}

func (cs *MessageHandlersSubrouter) handleTextQuestion(tctx telebot.Context) error {
	log.Println("called: handleTextQuestion")

	sender := tctx.Sender()

	data, ok := cs.KvClient.Get(strconv.Itoa(int(sender.ID)))

	if !ok {
		return tctx.Send("Не удалось получить данные :С")
	}

	text := tctx.Text()
	userQI := data.(*UserQuestionInfo)

	userQI.data = append(userQI.data, text)

	cs.KvClient.Set(strconv.Itoa(int(sender.ID)), &UserQuestionInfo{
		actionType: "handleTextAnswer",
		data:       userQI.data,
	})

	return tctx.Send("Введите ответ на вопрос:")
}

func (cs *MessageHandlersSubrouter) handleTextAnswer(tctx telebot.Context) error {
	log.Println("called: handleTextAnswer")

	ctx := context.Background()

	sender := tctx.Sender()

	kvData, ok := cs.KvClient.Get(strconv.Itoa(int(sender.ID)))

	if !ok {
		return tctx.Send("Не удалось получить данные :С")
	}

	text := tctx.Text()
	userQI := kvData.(*UserQuestionInfo)

	data := userQI.data[2:]

	answerValues := map[string]string{
		"type": "question",
		"data": text,
	}

	jsonValue, err := json.Marshal(answerValues)
	if err != nil {
		return tctx.Send("Не удалось получить данные :С")
	}

	conv, err := strconv.Atoi(data[0])
	if err != nil {
		return tctx.Send("Не удалось получить данные :С")
	}

	err = cs.QuestionsService.Add(ctx, int64(conv), data[1], string(jsonValue))

	if err != nil {
		return tctx.Send("Не удалось успешно записать вопрос :С")
	}

	return tctx.Send("Вопрос успешно записан!")
}
