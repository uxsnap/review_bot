package categoriesSubrouter

import (
	"context"
	"fmt"
	"log"
	"strings"

	"gopkg.in/telebot.v4"
)

func (cs *CategoriesSubrouter) addCategory(tctx telebot.Context) error {
	log.Println("called: addCategory")

	ctx := context.Background()

	args := tctx.Args()

	if len(args) < 2 {
		return tctx.Send(`
		Не хватает аргументов для добавления категории :С
Должны быть прокинуты ИМЯ ОПИСАНИЕ
	`)
	}

	name := strings.ToUpper(args[0])
	desc := strings.Join(args[1:], " ")

	err := cs.CategoriesService.Add(ctx, tctx.Update().Message.Sender.ID, name, desc)

	if err != nil {
		log.Printf("error: addCategory, %v", err)
		return tctx.Send("Не удалось добавить категорию :С")
	}

	count, err := cs.CategoriesService.Count(ctx, tctx.Update().Message.Sender.ID)

	if err != nil {
		log.Printf("error: addCategory, %v", err)
		return tctx.Send("Не удалось добавить категорию :С")
	}

	cs.KvClient.Set(
		fmt.Sprintf("%v_categories_count", tctx.Update().Message.Sender.ID),
		count,
	)

	log.Println("complete: addCategory")

	return tctx.Send(fmt.Sprintf("Добавлена новая категория: %v", name))
}
