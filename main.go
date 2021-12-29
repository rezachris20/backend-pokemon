package main

import (
	"log"
	"net/http"
	"pokemon-list/injector"
)

func main() {
	// e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	// }))

	// e.GET("/pokemon", controller.GetPokemon)
	// e.GET("/pokemon/:name", controller.DetailPokemon)

	// e.Logger.Fatal(e.Start(":1323"))
	// db := app.NewDB()
	// userRepository := repository.NewUserRepository(db)
	// userService := services.NewUserService(userRepository)
	// userController := controller.NewUserController(userService)
	// router := app.NewRouter(userController)
	// server := app.NewServer(router)
	s := injector.InitializedServer()
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
