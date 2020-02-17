package main

import (
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
	"log"
	"tgex/domain"
	"tgex/provider"

	"go.uber.org/fx"
	"tgex/internal/endpoint/tg"
)

func main() {
	app := fx.New(
		fx.NopLogger,

		fx.Provide(func() (*zap.Logger, error) {
			return zap.NewDevelopment()
		}),

		fx.Provide(func(logger *zap.Logger) (*domain.Config, error) {
			var conf domain.Config
			err := configor.Load(&conf, "config.toml")
			if err != nil {
				return nil, err
			}

			logger.Info("Load config", zap.Any("config", conf))

			return &conf, nil
		}),

		provider.Provider,

		fx.Invoke(tg.Register),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}
