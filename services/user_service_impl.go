package services

import (
	"errors"
	"pokemon-list/model"
	"pokemon-list/repository"
	"pokemon-list/web/request"
	"pokemon-list/web/response"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (s *UserServiceImpl) Register(payload *request.RegisterInput) (response.UserResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.MinCost)
	if err != nil {
		return response.UserResponse{}, err
	}

	user := &model.User{
		Nama:     payload.Nama,
		Username: payload.Username,
		Password: string(passwordHash),
	}

	registerUser, err := s.userRepository.Save(*user)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.ToUserResponse(registerUser), nil
}

func (s *UserServiceImpl) Login(payload *request.LoginInput) (response.UserResponse, error) {
	username := payload.Username
	password := payload.Password

	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return response.UserResponse{}, err
	}

	if user.ID == 0 {
		return response.UserResponse{}, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return response.UserResponse{}, errors.New("Password not match")
	}

	return response.ToUserResponse(user), nil
}
