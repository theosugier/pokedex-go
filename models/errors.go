package models

import "errors"

var ErrNotFound = errors.New("le pokémon n’existe pas")
var ErrBadRequest = errors.New("numéro de pokemon invalide")
var ErrTimeout = errors.New("API ne répond pas a temps")
var ErrUnavailable = errors.New("le service est indisponible")
