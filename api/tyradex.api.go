package api

import (
	"net/http"
	"time"

	"pokedex/models"
)

const baseURL = "https://tyradex.app/api/v1"

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func fetch(url string, cible any) error {
	// je prépare la requête
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "PokedexGo/1.0 (github.com/theosugier)")

	// j'appuie sur Send
	resp, err := client.Do(req)
	if err != nil {
		return models.ErrUnavailable
	}
	defer resp.Body.Close() // je referme la réponse à la fin

	return nil
}
