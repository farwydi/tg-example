package gateway

import "tgex/domain"

type ApodGateway interface {
	GetByDay(day string) (*domain.Apod, error)
}
