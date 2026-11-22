package main

import (
	"log"
	"net/http"

	"pokedex/helpers"
	"pokedex/routes"
)

func main() {
	helpers.ChargerTemplates()
	log.Println("serv lance sur http://localhost:8080")
	err := http.ListenAndServe(":8080", routes.Routeur())
	if err != nil {
		log.Fatal(err)
	}
}
