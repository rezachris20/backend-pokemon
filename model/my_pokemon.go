package model

import "time"

type MyPokemon struct {
	ID            int
	Name          string
	Nickname      string
	RealPokemonID int
	UserID        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
