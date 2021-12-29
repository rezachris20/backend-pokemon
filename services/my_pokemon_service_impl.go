package services

import (
	"fmt"
	"pokemon-list/helper"
	"pokemon-list/model"
	"pokemon-list/repository"
	"pokemon-list/web/request"
	"pokemon-list/web/response"
)

type MyPokemonServiceImpl struct {
	repository repository.MyPokemonRepository
}

func NewMyPokemonService(repository repository.MyPokemonRepository) MyPokemonService {
	return &MyPokemonServiceImpl{repository: repository}
}

func (m *MyPokemonServiceImpl) Register(payload *request.PokemonRegisterInput) (response.MyPokemonResponse, error) {
	tes := helper.GenerateFibonacci("Bula Basaur-9")
	fmt.Println(tes)

	/*@TODO
	- tambahin field update_count di table my_pokemon
	- ambil nilai terakhir update_count + 1
	- kirim hasil nya ke helper.GenerateFibonacci
	- concat str payload.Nickname + hasil dari helper.GenerateFibonacci
	*/
	return response.MyPokemonResponse{}, nil
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
