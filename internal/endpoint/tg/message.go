package tg

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (app App) Message(_ context.Context, update tgbotapi.Update) error {
	switch update.Message.Text {
	case "open":
		app.logger.Info("opne")
	}

	return nil
}
