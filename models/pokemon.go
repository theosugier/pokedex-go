package models

type Name struct {
	Nom string `json:"fr"`
}

type Sprites struct {
	Image string `json:"regular"`
}

type Type struct {
	Nom   string `json:"name"`
	Icone string `json:"image"`
}

type Pokemon struct {
	Numero    int     `json:"pokedex_id"`
	Categorie string  `json:"category"`
	Nom       Name    `json:"name"`
	Sprites   Sprites `json:"sprites"`
	Types     []Type  `json:"types"`
}