package handlers

import "github.com/labstack/echo/v4"

type LeadHandler interface {
	CreateLead(c echo.Context) error
}
