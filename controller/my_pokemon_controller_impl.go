package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"pokemon-list/helper"
	"pokemon-list/model"
	"pokemon-list/services"
	"pokemon-list/web/request"
	"strconv"
)

type MyPokemonControllerImpl struct {
	service services.MyPokemonService
}

func NewMyPokemonController(service services.MyPokemonService) MyPokemonController {
	return &MyPokemonControllerImpl{service: service}
}

func (m *MyPokemonControllerImpl) Register(c echo.Context) (err error) {
	input := new(request.PokemonRegisterInput)

	if err = c.Bind(input); err != nil {
		response := helper.APIResponse("Failed To Register", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if err = c.Validate(input); err != nil {
		response := helper.APIResponse("Register Failed", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	newPokemon, err := m.service.Register(input)
	if err != nil {
		response := helper.APIResponse("Failed To Register", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("Success To Register", http.StatusOK, "success", newPokemon)
	return c.JSON(http.StatusOK, response)
}

func (m *MyPokemonControllerImpl) Catch(c echo.Context) error {
	min := 10
	max := 100

	number := rand.Intn(max-min) + min
	if number < 50 {
		response := helper.APIResponse("Failed To Catch", http.StatusOK, "failed", nil)
		return c.JSON(http.StatusOK, response)
	}

	id := c.Param("id")
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+id, nil)
	if err != nil {
		response := helper.APIResponse("Failed To Catch", http.StatusBadRequest, "failed", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		response := helper.APIResponse("Failed To Catch", http.StatusBadRequest, "failed", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	defer resp.Body.Close()

	var data model.ResponseDetailPokemon

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		response := helper.APIResponse("Failed To Catch", http.StatusBadRequest, "failed", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success To Catch", http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, response)

}

func (m *MyPokemonControllerImpl) RenameNickName(c echo.Context) (err error) {
	input := new(request.PokemonRenameInput)
	if err = c.Bind(input); err != nil {
		response := helper.APIResponse("Failed To Rename", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if err = c.Validate(input); err != nil {
		response := helper.APIResponse("Register Failed", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	cekID, err := m.service.FindByID(id)
	if cekID.ID == 0 {
		response := helper.APIResponse("Failed To Rename", http.StatusNotFound, "false", nil)
		return c.JSON(http.StatusNotFound, response)
	}
	if err != nil {
		response := helper.APIResponse("Failed To Rename", http.StatusNotFound, "false", nil)
		return c.JSON(http.StatusNotFound, response)
	}
	renamePokemon, err := m.service.Rename(cekID.ID, input)
	if err != nil {
		response := helper.APIResponse("Failed To Rename", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("Success To Rename", http.StatusOK, "success", renamePokemon)
	return c.JSON(http.StatusOK, response)
}

func (m *MyPokemonControllerImpl) FindAllPokemonByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response := helper.APIResponse("Load Data Failed", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	pokemons, err := m.service.FindPokemonByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Load Data Failed", http.StatusBadRequest, "failed", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success To Load", http.StatusOK, "success", pokemons)
	return c.JSON(http.StatusOK, response)
}

func (m *MyPokemonControllerImpl) ReleasePokemon(c echo.Context) error {
	pokemonID, err := strconv.Atoi(c.Param("pokemon_id"))
	if err != nil {
		response := helper.APIResponse("Release Failed", http.StatusUnprocessableEntity, "failed", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	min := 1
	max := 100

	number := rand.Intn(max-min) + min
	isPrime := helper.CheckPrime(number)
	if !isPrime {
		response := helper.APIResponse("Release Failed", http.StatusUnprocessableEntity, "failed", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	_, err = m.service.DeletePokemon(pokemonID)
	if err != nil {
		response := helper.APIResponse("Release Failed", http.StatusUnprocessableEntity, "failed", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("Success Release", http.StatusOK, "success", nil)
	return c.JSON(http.StatusOK, response)
}
