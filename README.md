# Pokédex en Go

Site web qui affiche les Pokémon en français grâce à l'API Tyradex.

On peut voir la liste de tous les Pokémon, la fiche de chacun et chercher un Pokémon par son nom.

## Prérequis

- Go installé (https://go.dev/dl/)
- Une connexion internet

## Installation

```
git clone https://github.com/theosugier/pokedex-go.git
cd pokedex-go
```

## Lancement

```
go run .
```

Puis ouvrir http://localhost:8080 dans le navigateur.

## Exemples d'URL

- http://localhost:8080/pokemon : la liste de tous les Pokémon
- http://localhost:8080/pokemon/25 : la fiche de Pikachu
- http://localhost:8080/pokemon?nom=pika : la recherche par nom

## Source des données

Les données viennent de l'API Tyradex : https://tyradex.app

## Utilisation de l'IA

Tout le style du site (le fichier `assets/css/style.css`) a été fait avec une IA, comme le prof l'a autorisé.
