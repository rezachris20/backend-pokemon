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
	return c.JSON(http.StatusBadRequest, response)

}
