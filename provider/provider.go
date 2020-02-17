package provider

import (
	"errors"
	"go.uber.org/fx"
	"tgex/domain"
	"tgex/gateway/apodhttp"
	"tgex/usecase"
	"tgex/usecase/apod"
)

var Provider = fx.Options(
	fx.Provide(
		func(cfg *domain.Config) (usecase.ApodUseCase, error) {
			switch cfg.Provider.Apod {
			case "apodhttp":
				return apod.New(apodhttp.NewCache(cfg)), nil
			}

			return nil, errors.New("unknown apod provider")
		},
	),
)
