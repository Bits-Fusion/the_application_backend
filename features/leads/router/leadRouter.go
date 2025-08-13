package router

import (
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/permissions/entities"
	"github.com/labstack/echo/v4"

	leadHandlers "github.com/Bits-Fusion/the_application_backend/features/leads/handlers"
	leadRepo "github.com/Bits-Fusion/the_application_backend/features/leads/repositories"
	leadUsecase "github.com/Bits-Fusion/the_application_backend/features/leads/usecases"
)

type LeadRouter interface {
	Init(
		leadRouterGroupGroup *echo.Group,
		permissionMiddleware func(action entities.Action, resource string) echo.MiddlewareFunc,
	)
}

type router struct {
	db database.Database
}

func NewLeadRouter(db database.Database) *router {
	return &router{db: db}
}

func (r *router) Init(
	leadRouterGroup *echo.Group,
	permissionMiddleware func(action entities.Action, resource string) echo.MiddlewareFunc,
) {
	newLeadRepo := leadRepo.NewLeadRepository(r.db)
	newLeadUsecase := leadUsecase.NewLeadUsecase(newLeadRepo)
	newLeadHandler := leadHandlers.NewLeadHandler(newLeadUsecase)

	leadRouterGroup.POST("/", newLeadHandler.CreateLead, permissionMiddleware(entities.ActionCreate, "lead"))
	leadRouterGroup.GET("/", newLeadHandler.ListLeads, permissionMiddleware(entities.ActionView, "lead"))
	leadRouterGroup.PATCH("/:leadId", newLeadHandler.UpdateLead, permissionMiddleware(entities.ActionUpdate, "lead"))
	leadRouterGroup.DELETE("/:leadId", newLeadHandler.DeleteLead, permissionMiddleware(entities.ActionDelete, "lead"))
}
