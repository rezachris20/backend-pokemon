package request

type PokemonRegisterInput struct {
	PokemonID   int    `json:"pokemon_id" validate:"required"`
	PokemonName string `json:"pokemon_name" validate:"required"`
	NickName    string `json:"nick_name" validate:"required"`
	UserID      int    `json:"user_id" validate:"required"`
}

type PokemonRenameInput struct {
	NickName string `json:"nick_name" validate:"required"`
}
