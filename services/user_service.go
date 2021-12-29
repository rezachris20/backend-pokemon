package services

import (
	"pokemon-list/web/request"
	"pokemon-list/web/response"
)

type UserService interface {
	Register(payload *request.RegisterInput) (response.UserResponse, error)
	Login(payload *request.LoginInput) (response.UserResponse, error)
}
