package controllers

import (
	"errors"
	"log"
	"net/http"

	"pokedex/helpers"
	"pokedex/models"
	"pokedex/services"
)

func ListePokemon(w http.ResponseWriter, r *http.Request) {
	liste, err := services.GetAllPokemon()
	if err != nil {
		afficherErreur(w, err)
		return
	}
	donnees := map[string]any{"Pokemons": liste}             // je range la liste avk nom "Pokemons" pour la page
	helpers.Afficher(w, "list.html", http.StatusOK, donnees) // je remplis list.html avec la liste & jl'envoie au visiteur
}

func FichePokemon(w http.ResponseWriter, r *http.Request) {
	p, err := services.GetPokemon(r.PathValue("id"))
	if err != nil {
		afficherErreur(w, err)
		return
	}
	helpers.Afficher(w, "pokemon.html", http.StatusOK, p)
}

func afficherErreur(w http.ResponseWriter, err error) {
	log.Println("erreur :", err)
	code := http.StatusServiceUnavailable
	message := "Service momentanément indisponible."
	if errors.Is(err, models.ErrBadRequest) {
		code = http.StatusBadRequest
		message = "Numéro de Pokémon invalide."
	} else if errors.Is(err, models.ErrNotFound) {
		code = http.StatusNotFound
		message = "Ce Pokémon n'existe pas."
	}
	donnees := map[string]any{"Code": code, "Message": message}
	helpers.Afficher(w, "error.html", code, donnees)
}

func PageInconnue(w http.ResponseWriter, r *http.Request) {
	donnees := map[string]any{"Code": http.StatusNotFound, "Message": "Cette page n'existe pas."}
	helpers.Afficher(w, "error.html", http.StatusNotFound, donnees)
}
