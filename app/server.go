package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewServer(router *echo.Echo) *http.Server {
	return &http.Server{Addr: ":1324", Handler: router}
}
