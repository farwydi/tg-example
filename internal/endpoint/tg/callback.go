package tg

import (
	"context"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	"tgex/internal/endpoint/tg/menu"
)

func (app App) Callback(_ context.Context, update tgbotapi.Update) error {
	splitData := strings.Split(update.CallbackQuery.Data, " ")
	if len(splitData) < 2 {
		return errors.New("bad state")
	}
	switch splitData[0] {
	case "next":
		apod, err := app.ApodUseCase.GetOneDay(splitData[1])
		if err != nil {
			return err
		}

		msg := tgbotapi.NewEditMessageText(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			fmt.Sprintf("*%s*\n\n%s\n%s",
				apod.Title, apod.Explanation, apod.Image,
			),
		)
		msg.ParseMode = "markdown"
		*msg.ReplyMarkup = menu.ApodMenuBuild(app.ApodUseCase.CallbackForm(apod.Date))

		_, err = app.bot.Send(msg)
		if err != nil {
			return err
		}
	case "prev":
		apod, err := app.ApodUseCase.GetOneDay(splitData[1])
		if err != nil {
			return err
		}

		_, err = app.bot.Send(tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{
				ChatID:    update.CallbackQuery.Message.Chat.ID,
				MessageID: update.CallbackQuery.Message.MessageID,
			},
			ParseMode: "markdown",
			Text: fmt.Sprintf("*%s*\n\n%s\n%s",
				apod.Title, apod.Explanation, apod.Image,
			),
		})
		if err != nil {
			return err
		}

		_, err = app.bot.Send(tgbotapi.NewEditMessageReplyMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			menu.ApodMenuBuild(app.ApodUseCase.CallbackForm(apod.Date)),
		))
		if err != nil {
			return err
		}
	}

	return nil
}
