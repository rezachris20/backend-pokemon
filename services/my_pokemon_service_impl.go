package services

import (
	"pokemon-list/helper"
	"pokemon-list/model"
	"pokemon-list/repository"
	"pokemon-list/web/request"
	"pokemon-list/web/response"
	"strconv"
)

type MyPokemonServiceImpl struct {
	repository repository.MyPokemonRepository
}

func NewMyPokemonService(repository repository.MyPokemonRepository) MyPokemonService {
	return &MyPokemonServiceImpl{repository: repository}
}

func (m *MyPokemonServiceImpl) Register(payload *request.PokemonRegisterInput) (response.MyPokemonResponse, error) {
	myPokemon := &model.MyPokemon{
		Name:          payload.PokemonName,
		Nickname:      payload.NickName + "-0",
		RealPokemonID: payload.PokemonID,
		UserID:        payload.UserID,
	}

	save, err := m.repository.Save(*myPokemon)
	if err != nil {
		return response.MyPokemonResponse{}, err
	}

	return response.ToMyPokemonResponse(save), nil
}

func (m *MyPokemonServiceImpl) Rename(ID int, payload *request.PokemonRenameInput) (response.MyPokemonResponse, error) {
	pokemon, err := m.repository.FindByID(ID)
	if err != nil {
		return response.ToMyPokemonResponse(pokemon), err
	}

	generateFibonacci := helper.FibonacciRecursion(pokemon.Counter)
	number := strconv.Itoa(generateFibonacci)

	pokemon.Nickname = payload.NickName + "-" + number
	pokemon.Counter = pokemon.Counter + 1

	edit, err := m.repository.Update(pokemon.ID, pokemon)
	if err != nil {
		return response.ToMyPokemonResponse(edit), err
	}

	return response.ToMyPokemonResponse(edit), nil
}

func (m *MyPokemonServiceImpl) FindByID(ID int) (response.MyPokemonResponse, error) {
	pokemon, err := m.repository.FindByID(ID)
	if err != nil {
		return response.ToMyPokemonResponse(pokemon), err
	}
	return response.ToMyPokemonResponse(pokemon), nil
}
