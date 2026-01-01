package internal

import (
	"net/http"
	"time"
)

const (
	baseUrl = "https://opensky-network.org/api"
)

func NewClient() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
}
