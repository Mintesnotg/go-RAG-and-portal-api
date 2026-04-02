package routes

import (
	"go-api/internal/handlers"
	"go-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(router *gin.Engine, roleHandler *handlers.RoleHandler) {
	rolesGroup := router.Group("/api/roles")
	rolesGroup.Use(middleware.AuthMiddleware())
	rolesGroup.POST("/assign", roleHandler.AssignRole)
}
