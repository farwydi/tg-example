package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
	"time"
)

func (app App) Setup() error {
	_, err := app.bot.SetWebhook(
		tgbotapi.NewWebhook(app.endpoint + app.bot.Token),
	)
	if err != nil {
		return err
	}

	info, err := app.bot.GetWebhookInfo()
	if err != nil {
		return err
	}

	lastErrorTime := time.Unix(int64(info.LastErrorDate), 0)
	if lastErrorTime.After(app.start) {
		app.logger.Error("Telegram callback failed", zap.String("info", info.LastErrorMessage))
	}

	return nil
}
