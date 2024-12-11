package categoriesSubrouter

import (
	"context"
	"fmt"
	"log"
	"strings"

	"gopkg.in/telebot.v4"
)

func (cs *CategoriesSubrouter) deleteCategory(tctx telebot.Context) error {
	log.Println("called: deleteCategory")

	ctx := context.Background()

	args := tctx.Args()

	if len(args) != 1 {
		return tctx.Send(`
		Не хватает аргументов для удаления категории :С
Должно быть прокинуто ИМЯ
	`)
	}

	name := strings.ToUpper(args[0])

	err := cs.CategoriesService.Del(ctx, tctx.Update().Message.Sender.ID, name)

	if err != nil {
		log.Printf("error: deleteCategory, %v", err)
		return tctx.Send("Не удалось удалить категорию :С")
	}

	log.Println("complete: deleteCategory")

	return tctx.Send(fmt.Sprintf("Удалена категория: %v", name))
}
