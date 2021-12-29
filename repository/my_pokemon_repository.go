package repository

import "pokemon-list/model"

type MyPokemonRepository interface {
	Save(myPokemon model.MyPokemon) (model.MyPokemon, error)
}
