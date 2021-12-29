package app

import (
	"net/http"
	"pokemon-list/controller"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRouter(userController controller.UserController, pokemonController controller.MyPokemonController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", userController.Login)
	e.POST("/register", userController.Register)

	e.GET("/pokemon", controller.GetPokemon)
	e.GET("/pokemon/:name", controller.DetailPokemon)

	e.POST("/pokemon/catch/:id", pokemonController.Catch)
	e.POST("/pokemon/register", pokemonController.Register)
	e.POST("/pokemon/rename/:id", pokemonController.RenameNickName)
	return e
}
