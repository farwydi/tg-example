package apodhttp

import (
	"github.com/patrickmn/go-cache"
	"tgex/domain"
	"tgex/gateway"
	"time"
)

func NewCache(cfg *domain.Config) gateway.ApodGateway {
	return &cacheGateway{
		ApodGateway: New(cfg),
		Cache:       cache.New(5*time.Minute, 10*time.Minute),
	}
}

type cacheGateway struct {
	gateway.ApodGateway
	Cache *cache.Cache
}

func (gw cacheGateway) GetByDay(day string) (*domain.Apod, error) {
	if apod, found := gw.Cache.Get(day); found {
		return apod.(*domain.Apod), nil
	}

	apod, err := gw.ApodGateway.GetByDay(day)
	if err != nil {
		return nil, err
	}

	gw.Cache.Set(day, apod, cache.DefaultExpiration)
	return apod, nil
}
