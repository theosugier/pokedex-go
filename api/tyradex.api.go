package api

import (
	"net/http"
	"time"
)

const baseURL = "https://tyradex.app/api/v1"

var client = &http.Client{
	Timeout: 5 * time.Second,
}
