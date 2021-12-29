package controller

import (
	"net/http"
	"pokemon-list/helper"
	"pokemon-list/services"
	"pokemon-list/web/request"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	userServices services.UserService
}

func NewUserController(userServices services.UserService) UserController {
	return &UserControllerImpl{userServices: userServices}
}

func (h *UserControllerImpl) Register(c echo.Context) (err error) {

	input := new(request.RegisterInput)

	if err = c.Bind(input); err != nil {
		response := helper.APIResponse("Register Failed", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if err = c.Validate(input); err != nil {
		response := helper.APIResponse("Register Failed", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	register, err := h.userServices.Register(input)
	if err != nil {
		response := helper.APIResponse("Register Failed", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("Account has been registerd", http.StatusCreated, "success", register)
	return c.JSON(http.StatusCreated, response)

}

func (h *UserControllerImpl) Login(c echo.Context) (err error) {
	input := new(request.LoginInput)

	if err = c.Bind(input); err != nil {
		response := helper.APIResponse("Register Failed", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if err = c.Validate(input); err != nil {
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "failed", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	loggin, err := h.userServices.Login(input)
	if err != nil {
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "failed", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Successfully loggedin", http.StatusOK, "success", loggin)
	return c.JSON(http.StatusOK, response)
}
