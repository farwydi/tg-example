package apodhttp

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"strings"
	"tgex/domain"
	"tgex/gateway"
	"time"
)

const _apodUrl = "https://api.nasa.gov/planetary/apod?api_key="

func New(cfg *domain.Config) gateway.ApodGateway {
	return &apodGateway{
		client: resty.New(),
		apiKey: cfg.NasaApiKey,
	}
}

type apodGateway struct {
	client *resty.Client
	apiKey string
}

func escapedStr(t string) string {
	t = strings.ReplaceAll(t, "_", "\\_")
	t = strings.ReplaceAll(t, "*", "\\*")
	t = strings.ReplaceAll(t, "[", "\\[")
	t = strings.ReplaceAll(t, "`", "\\`")
	return t
}

func (gw apodGateway) GetByDay(day string) (*domain.Apod, error) {
	resp, err := gw.client.R().
		SetQueryParam("date", day).
		Get(_apodUrl + gw.apiKey)
	if err != nil {
		return nil, err
	}

	var result Apod
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	dataTime, err := time.Parse("2006-01-02", result.Date)
	if err != nil {
		return nil, err
	}

	return &domain.Apod{
		Date:           dataTime,
		Explanation:    result.Explanation,
		ServiceVersion: result.ServiceVersion,
		Title:          result.Title,
		Image:          escapedStr(result.URL),
	}, nil
}
