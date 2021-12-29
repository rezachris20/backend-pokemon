package repository

import (
	"gorm.io/gorm"
	"pokemon-list/model"
)

type MyPokemonRepositoryImpl struct {
	db *gorm.DB
}

func NewMyPokemonRepository(db *gorm.DB) MyPokemonRepository {
	return &MyPokemonRepositoryImpl{db: db}
}

func (m *MyPokemonRepositoryImpl) Save(myPokemon model.MyPokemon) (model.MyPokemon, error) {
	err := m.db.Create(&myPokemon).Error
	if err != nil {
		return myPokemon, err
	}
	return myPokemon, nil
}
