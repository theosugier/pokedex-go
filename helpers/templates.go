package helpers

import (
	"html/template"
	"log"
	"net/http"
)

var pages *template.Template

func ChargerTemplates() {
	pages = template.Must(template.ParseGlob("templates/*.html"))
}

func Afficher(w http.ResponseWriter, nom string, code int, donnees any) {
	w.WriteHeader(code)
	err := pages.ExecuteTemplate(w, nom, donnees)
	if err != nil {
		log.Println("erreur template :", err)
	}
}
