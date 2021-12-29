//go:build wireinject
// +build wireinject

package injector

import (
	"net/http"
	"pokemon-list/app"
	"pokemon-list/controller"
	"pokemon-list/repository"
	"pokemon-list/services"

	"github.com/google/wire"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	services.NewUserService,
	controller.NewUserController,
)

var myPokemonSet = wire.NewSet(
	repository.NewMyPokemonRepository,
	services.NewMyPokemonService,
	controller.NewMyPokemonController,
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		userSet,
		myPokemonSet,
		app.NewRouter,
		app.NewServer,
	)
	return nil
}
