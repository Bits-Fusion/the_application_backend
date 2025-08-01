package handlers

import "github.com/labstack/echo/v4"

type UserHandler interface {
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
	// ForgotPassword(c echo.Context) error

	ListUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	GetUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}
