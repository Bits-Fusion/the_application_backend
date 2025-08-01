package server

import (
	// "fmt"
	"log"
	"net/http"
	"strings"

	permissionEntity "github.com/Bits-Fusion/the_application_backend/features/permissions/entities"
	"github.com/Bits-Fusion/the_application_backend/features/users/entities"
	userRepo "github.com/Bits-Fusion/the_application_backend/features/users/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/gommon/log"
)

func (server *echoServer) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Authorization header"})
		}

		token := parts[1]
		jwtToken, err := server.auth.ValidateToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		// userID := fmt.Sprintf("%.f", claims["sub"])
		c.Set("user_id", claims["sub"])

		return next(c)
	}
}

func (s *echoServer) RequirePermission(action permissionEntity.Action, resource string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userId := c.Get("user_id").(string)
			uid, err := uuid.Parse(userId)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
			}

			userRepository := userRepo.NewUserPostgresRepository(s.db)
			user, err := userRepository.GetUserData(entities.FilterByID, uid.String())

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			log.Print(user)
			for _, perm := range user.Permissions {
				if perm.Action == action && perm.Resource == resource {
					return next(c)
				}
			}
			return echo.NewHTTPError(http.StatusForbidden, "this user does not have this pemission")
		}
	}
}
