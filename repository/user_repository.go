package repository

import "pokemon-list/model"

type UserRepository interface {
	Save(user model.User) (model.User, error)
	FindByUsername(username string) (model.User, error)
}
