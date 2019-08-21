package categories

import (
	"github.com/gin-gonic/gin"
	"ngtium/libraries/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	categories := r.Group("/categories")
	{
		categories.POST("/:project", middlewares.Authorized, middlewares.Project, create)
		categories.GET("/:category/:action", middlewares.Category, list)
		categories.DELETE("/:category/:project", middlewares.Authorized, remove)
		categories.PATCH("/:category/:project", middlewares.Authorized, middlewares.Category, middlewares.Project, update)
	}
}
