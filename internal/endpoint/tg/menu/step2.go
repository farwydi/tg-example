package menu

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var Step2Menu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("step 3", "GoToFuck 2"),
	),
)
