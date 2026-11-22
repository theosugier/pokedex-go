package routes

import (
	"net/http"

	"pokedex/controllers"
)

func Routeur() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /pokemon", controllers.ListePokemon)
	mux.HandleFunc("GET /pokemon/{id}", controllers.FichePokemon)
	mux.Handle("GET /{$}", http.RedirectHandler("/pokemon", http.StatusMovedPermanently))
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	return mux
}
