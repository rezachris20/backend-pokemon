package main

import (
	"pokemon-list/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/pokemon", controller.GetPokemon)
	e.GET("/pokemon/:name", controller.DetailPokemon)
	e.Logger.Fatal(e.Start(":1323"))
}
