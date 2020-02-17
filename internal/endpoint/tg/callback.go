package tg

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"regexp"
	"tgex/internal/endpoint/tg/menu"
)

var _prevReg = regexp.MustCompile(`prev_(.*)`)
var _nextReg = regexp.MustCompile(`next_(.*)`)

func (app App) Callback(_ context.Context, update tgbotapi.Update) error {
	switch {
	case _nextReg.MatchString(update.CallbackQuery.Data):
		date := _nextReg.FindAllStringSubmatch(update.CallbackQuery.Data, 1)
		apod, err := app.ApodUseCase.GetOneDay(date[0][1])
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
	case _prevReg.MatchString(update.CallbackQuery.Data):
		date := _prevReg.FindAllStringSubmatch(update.CallbackQuery.Data, 1)
		apod, err := app.ApodUseCase.GetOneDay(date[0][1])
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
