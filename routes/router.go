package routes

import (
	"go-backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(roleHandler *handlers.RoleHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/roles", roleHandler.CreateRole)
		api.PATCH("/roles/:id", roleHandler.UpdateRole)
		api.DELETE("/roles/:id", roleHandler.DeleteRole)
		api.GET("/roles", roleHandler.GetRoles)
		api.GET("/roles/:id", roleHandler.GetRoleByID)
	}

	return r
}
