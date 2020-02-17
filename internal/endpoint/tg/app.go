package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"tgex/usecase"
	"time"
)

var _app *App

type Params struct {
	fx.In

	ApodUseCase usecase.ApodUseCase
}

type App struct {
	Params
	bot      *tgbotapi.BotAPI
	logger   *zap.Logger
	updates  tgbotapi.UpdatesChannel
	start    time.Time
	done     chan struct{}
	endpoint string
}

func (app *App) Init(token string) error {
	err := tgbotapi.SetLogger(NewLogger(app.logger))
	if err != nil {
		return err
	}

	app.bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	//app.bot.Debug = true

	app.logger.Info("Authorized on account", zap.String("name", app.bot.Self.UserName))

	app.updates = app.bot.ListenForWebhook("/" + app.bot.Token)

	return nil
}
