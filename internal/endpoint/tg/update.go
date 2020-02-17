package tg

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (app App) Update(ctx context.Context, update tgbotapi.Update) error {
	if update.CallbackQuery != nil {
		err := app.Callback(ctx, update)
		if err != nil {
			return err
		}

		return nil
	}

	if update.Message != nil {
		if update.Message.IsCommand() {
			err := app.Command(ctx, update)
			if err != nil {
				return err
			}

			return nil
		}

		err := app.Message(ctx, update)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}
