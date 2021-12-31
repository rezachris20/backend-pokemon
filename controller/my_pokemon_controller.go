package controller

import "github.com/labstack/echo/v4"

type MyPokemonController interface {
	Catch(c echo.Context) error
	Register(c echo.Context) error
	RenameNickName(c echo.Context) error
	FindAllPokemonByUserID(c echo.Context) error
	ReleasePokemon(c echo.Context) error
}
