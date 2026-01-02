package internal

import (
	"log"
	"net/http"
	"time"

	"github.com/matthieukhl/go-atc/internal/config"
)

const (
	baseUrl = "https://opensky-network.org/api"
)

type Client struct {
	Config     config.Config
	HTTPClient http.Client
}

func NewClient() Client {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	return Client{
		Config: cfg,
		HTTPClient: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
