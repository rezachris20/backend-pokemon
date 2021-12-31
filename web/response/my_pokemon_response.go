package response

import "pokemon-list/model"

type MyPokemonResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	NickName      string `json:"nick_name"`
	RealPokemonID int    `json:"real_pokemon_id"`
	UserID        int    `json:"user_id"`
	ImageURL      string `json:"image_url"`
}

func ToMyPokemonResponse(pokemon model.MyPokemon) MyPokemonResponse {
	return MyPokemonResponse{
		ID:            pokemon.ID,
		Name:          pokemon.Name,
		NickName:      pokemon.Nickname,
		RealPokemonID: pokemon.RealPokemonID,
		UserID:        pokemon.UserID,
		ImageURL:      pokemon.ImageURL,
	}
}

func ToMyPokemonResponses(pokemons []model.MyPokemon) []MyPokemonResponse {
	var pokemonResponses []MyPokemonResponse
	for _, pokemon := range pokemons {
		pokemonResponses = append(pokemonResponses, ToMyPokemonResponse(pokemon))
	}
	return pokemonResponses
}
