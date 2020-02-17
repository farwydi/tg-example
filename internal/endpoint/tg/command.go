package tg

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"tgex/internal/endpoint/tg/menu"
)

func (app App) Command(_ context.Context, update tgbotapi.Update) error {
	switch update.Message.Command() {
	case "start":
		apod, err := app.ApodUseCase.GetOneDay(app.ApodUseCase.Now())
		if err != nil {
			return err
		}

		day, next, prev := app.ApodUseCase.CallbackForm(apod.Date)
		msg := tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID:      update.Message.Chat.ID,
				ReplyMarkup: menu.ApodMenuBuild(day, next, prev),
			},
			Text: fmt.Sprintf("*%s*\n\n%s\n%s",
				apod.Title, apod.Explanation, apod.Image,
			),
			ParseMode: "markdown",
		}
		_, err = app.bot.Send(msg)
		if err != nil {
			return err
		}
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command")
		_, err := app.bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
