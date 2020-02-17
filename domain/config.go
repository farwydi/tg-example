package domain

type ProviderConfig struct {
	Apod string
}

type Config struct {
	NasaApiKey string `env:"NASA_API_KEY"`
	Endpoint   string `env:"ENDPOINT_URL"`
	Provider   ProviderConfig
}
