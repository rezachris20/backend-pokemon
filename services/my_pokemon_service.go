package services

import (
	"pokemon-list/web/request"
	"pokemon-list/web/response"
)

type MyPokemonService interface {
	Register(payload *request.PokemonRegisterInput) (response.MyPokemonResponse, error)
}
