package menu

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ApodMenuBuild(isToday bool, next, prev string) tgbotapi.InlineKeyboardMarkup {
	if isToday {
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Назад: "+prev, "prev "+prev),
			),
		)
	}

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад: "+prev, "prev "+prev),
			tgbotapi.NewInlineKeyboardButtonData("Дальше: "+next, "next "+next),
		),
	)
}
