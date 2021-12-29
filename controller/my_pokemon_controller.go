package controller

import "github.com/labstack/echo/v4"

type MyPokemonController interface {
	Catch(c echo.Context) error
	Register(c echo.Context) error
}
