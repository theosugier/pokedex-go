package controllers

import (
	"log"
	"net/http"

	"pokedex/helpers"
	"pokedex/services"
)

func ListePokemon(w http.ResponseWriter, r *http.Request) {
	liste, err := services.GetAllPokemon()
	if err != nil {
		log.Println("erreur liste :", err)
		http.Error(w, "Service indisponible.", http.StatusServiceUnavailable)
		return
	}
	donnees := map[string]any{"Pokemons": liste}             // je range la liste avk nom "Pokemons" pour la page
	helpers.Afficher(w, "list.html", http.StatusOK, donnees) // je remplis list.html avec la liste & jl'envoie au visiteur
}

func FichePokemon(w http.ResponseWriter, r *http.Request) {
	p, err := services.GetPokemon(r.PathValue("id"))
	if err != nil {
		log.Println("erreur fiche :", err)
		http.Error(w, "Impossible d'afficher ce Pokémon.", http.StatusNotFound)
		return
	}
	helpers.Afficher(w, "pokemon.html", http.StatusOK, p)
}
