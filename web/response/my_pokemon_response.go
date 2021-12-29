package response

import "pokemon-list/model"

type MyPokemonResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	NickName      string `json:"nick_name"`
	RealPokemonID int    `json:"real_pokemon_id"`
	UserID        int    `json:"user_id"`
}

func ToMyPokemonResponse(pokemon model.MyPokemon) MyPokemonResponse {
	return MyPokemonResponse{
		ID:            pokemon.ID,
		Name:          pokemon.Name,
		NickName:      pokemon.Nickname,
		RealPokemonID: pokemon.RealPokemonID,
		UserID:        pokemon.UserID,
	}
}
