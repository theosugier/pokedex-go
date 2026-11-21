package api

import (
	"encoding/json"
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

	// je regarde le code avant d'ouvrir
	if resp.StatusCode == http.StatusNotFound {
		return models.ErrNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return models.ErrUnavailable
	}

	// je range le JSON dans la boîte
	return json.NewDecoder(resp.Body).Decode(cible)
}

func GetAllPokemon() ([]models.Pokemon, error) {
	var liste []models.Pokemon
	err := fetch(baseURL+"/pokemon", &liste)
	return liste, err
}
