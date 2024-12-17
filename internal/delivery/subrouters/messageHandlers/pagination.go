package messageHandlersSubrouter

import (
	"fmt"

	"github.com/uxsnap/review_bot/internal/delivery/subrouters"
	"gopkg.in/telebot.v4"
)

type PaginationData struct {
	selector    *telebot.ReplyMarkup
	curPageInt  int
	endpoint    string
	maxSizeName string
}

func (cs *MessageHandlersSubrouter) Pagination(tctx telebot.Context, pData PaginationData) []telebot.Btn {
	buttonRows := []telebot.Btn{}

	if pData.curPageInt != 0 {
		buttonRows = append(buttonRows,
			pData.selector.Data("⬅", pData.endpoint, "button_prev", fmt.Sprintf("%v", pData.curPageInt-1)),
		)
	}

	categoriesCountObject, kvOk := cs.KvClient.Get(
		fmt.Sprintf("%v_%v", tctx.Update().Message.Sender.ID, pData.maxSizeName),
	)

	categoriesCount, typeCaseOk := categoriesCountObject.(int)

	if kvOk && typeCaseOk {
		if categoriesCount > subrouters.LIMIT_COUNT*pData.curPageInt {
			buttonRows = append(
				buttonRows, pData.selector.Data("➡", pData.endpoint, "button_next", fmt.Sprintf("%v", pData.curPageInt+1)),
			)
		}
	}

	return buttonRows
}
