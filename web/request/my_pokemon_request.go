package request

type PokemonRegisterInput struct {
	PokemonID   int    `json:"pokemon_id"`
	PokemonName string `json:"pokemon_name"`
	NickName    string `json:"nick_name"`
	UserID      int    `json:"user_id"`
}
