package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokemon-list/entity"

	"github.com/labstack/echo/v4"
)

func GetPokemon(c echo.Context) error {

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon", nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer resp.Body.Close()

	var data entity.ResponsePokemon

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, data)
}

func DetailPokemon(c echo.Context) error {
	name := c.Param("name")

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+name, nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer resp.Body.Close()

	var data entity.ResponseDetailPokemon

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, data)

}
