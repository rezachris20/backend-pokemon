package repository

import (
	"fmt"
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

func (m *MyPokemonRepositoryImpl) Update(ID int, myPokemon model.MyPokemon) (model.MyPokemon, error) {
	fmt.Println(ID)
	err := m.db.Where("id = ?", ID).Save(&myPokemon).Error
	if err != nil {
		return myPokemon, err
	}
	return myPokemon, nil
}

func (m *MyPokemonRepositoryImpl) FindByID(ID int) (model.MyPokemon, error) {
	var pokemon model.MyPokemon
	err := m.db.Where("id = ? ", ID).Find(&pokemon).Error
	if err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func (m *MyPokemonRepositoryImpl) FindByUserID(ID int) ([]model.MyPokemon, error) {
	var pokemons []model.MyPokemon
	err := m.db.Where("user_id = ?", ID).Find(&pokemons).Error
	if err != nil {
		return pokemons, err
	}

	return pokemons, nil
}

func (m *MyPokemonRepositoryImpl) Delete(ID int) (bool, error) {
	var pokemon model.MyPokemon
	err := m.db.Where("id = ?", ID).Delete(&pokemon)
	if err != nil {
		return false, nil
	}
	return true, nil
}
