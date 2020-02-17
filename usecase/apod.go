package usecase

import (
	"tgex/domain"
	"time"
)

type ApodUseCase interface {
	GetOneDay(day string) (apod *domain.Apod, err error)
	CallbackForm(day time.Time) (bool, string, string)
	IsToday(day time.Time) bool
	Next(day time.Time) string
	Prev(day time.Time) string
	Now() string
}
