package categoriesSubrouter

import (
	"context"
	"fmt"
	"log"
	"strings"

	"gopkg.in/telebot.v4"
)

func (cs *CategoriesSubrouter) getAllCategories(tctx telebot.Context) error {
	log.Println("call: getAllCategories")

	ctx := context.Background()

	var name string
	args := tctx.Args()

	if len(args) != 0 {
		name = args[0]
	}

	categories, err := cs.CategoriesService.Get(ctx, tctx.Update().Message.Sender.ID, name, -1, -1)

	if err != nil {
		log.Printf("error: getAllCategories, %v", err)
		return tctx.Send("Не найдено таких категорий :С")
	}

	categoryRes := make([]string, len(categories))

	for ind, c := range categories {
		categoryRes = append(categoryRes, fmt.Sprintf(
			"\n%v. Название: %v,\nОписание: %v", ind, c.Name, c.Description,
		))
	}

	log.Println("complete: getAllCategories")

	if len(categories) == 0 {
		return tctx.Send("Категорий пока нет...")
	}

	return tctx.Send(strings.Join(categoryRes, "\n"))
}
