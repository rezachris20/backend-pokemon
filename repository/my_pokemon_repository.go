package repository

import "pokemon-list/model"

type MyPokemonRepository interface {
	Save(myPokemon model.MyPokemon) (model.MyPokemon, error)
	Update(ID int, myPokemon model.MyPokemon) (model.MyPokemon, error)
	FindByID(ID int) (model.MyPokemon, error)
}
