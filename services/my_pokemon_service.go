package services

import (
	"pokemon-list/web/request"
	"pokemon-list/web/response"
)

type MyPokemonService interface {
	Register(payload *request.PokemonRegisterInput) (response.MyPokemonResponse, error)
	Rename(ID int, payload *request.PokemonRenameInput) (response.MyPokemonResponse, error)
	FindByID(ID int) (response.MyPokemonResponse, error)
}
