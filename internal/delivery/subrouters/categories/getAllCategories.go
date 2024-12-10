package categoriesSubrouter

import (
	"context"
	"log"

	"gopkg.in/telebot.v4"
)

func (cs *CategoriesSubrouter) getAllCategories(tctx telebot.Context) error {
	log.Println("getAllCategories")

	ctx := context.Background()

	var name string
	args := tctx.Args()

	if len(args) != 0 {
		name = args[0]
	}

	categories, err := cs.CategoriesService.Get(ctx, name)

	if err != nil {
		log.Printf("error: getAllCategories, %v", err)
		return tctx.Send("Не найдено таких категорий :С")
	}

	categoryRes := make([]string, len(categories))

	for _, c := range categories {
		categoryRes = append(categoryRes, "Название: %v,\nОписание: %v", c.Name, c.Description)
	}

	log.Println("complete: getAllCategories")

	return tctx.Send(categoryRes)
}
