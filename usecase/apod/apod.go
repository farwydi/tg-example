package apod

import (
	"tgex/domain"
	"tgex/gateway"
	"tgex/usecase"
	"time"
)

const _dayForm = "2006-01-02"

func New(gw gateway.ApodGateway) usecase.ApodUseCase {
	return &apodUseCase{
		ApodGateway: gw,
	}
}

type apodUseCase struct {
	ApodGateway gateway.ApodGateway
}

func (uc apodUseCase) Now() string {
	return time.Now().Format(_dayForm)
}

func (uc apodUseCase) CallbackForm(day time.Time) (bool, string, string) {
	return uc.IsToday(day), uc.Next(day), uc.Prev(day)
}

func (uc apodUseCase) Next(day time.Time) string {
	return day.AddDate(0, 0, 1).
		Format(_dayForm)
}

func (uc apodUseCase) Prev(day time.Time) string {
	return day.AddDate(0, 0, -1).
		Format(_dayForm)
}

func (uc apodUseCase) IsToday(day time.Time) bool {
	y, m, d := day.Date()
	yn, mn, dn := time.Now().Date()
	return yn == y && mn == m && dn == d
}

func (uc apodUseCase) GetOneDay(day string) (*domain.Apod, error) {
	apod, err := uc.ApodGateway.GetByDay(day)
	if err != nil {
		return nil, err
	}

	return apod, nil
}
