package handlers

import "github.com/labstack/echo/v4"

type UserHandler interface {
	SignUp(c echo.Context) error
}
