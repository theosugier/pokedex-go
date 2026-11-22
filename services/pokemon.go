package services

import (
	"pokedex/api"
	"pokedex/models"
)

func GetAllPokemon() ([]models.Pokemon, error) {
	liste, err := api.GetAllPokemon()
	if err != nil {
		return nil, err
	}
	var vrais []models.Pokemon
	for _, p := range liste {
		if p.Numero > 0 {
			vrais = append(vrais, p)
		}
	}
	return vrais, nil
}
