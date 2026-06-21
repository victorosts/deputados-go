package camara

import "time"

type Config struct {
	BaseURL string
	Timeout time.Duration
}

func DefaultConfig() Config {
	return Config{
		BaseURL: "https://dadosabertos.camara.leg.br/api/v2/",
		Timeout: 10 * time.Second,
	}
}
