package handlers

import "github.com/labstack/echo/v4"

type LeadHandler interface {
	CreateLead(c echo.Context) error
	ListLeads(c echo.Context) error
	UpdateLead(c echo.Context) error
	DeleteLead(c echo.Context) error
}
