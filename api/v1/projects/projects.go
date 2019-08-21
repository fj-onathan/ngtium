package projects

import (
	"github.com/gin-gonic/gin"
	"ngtium/libraries/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	projects := r.Group("/projects")
	{
		projects.POST("/", middlewares.Authorized, create)
		projects.GET("/", list)
		projects.GET("/:project", middlewares.Authorized, middlewares.Project, read)
		projects.DELETE("/:project", middlewares.Authorized, middlewares.Project, remove)
		projects.PATCH("/:project", middlewares.Authorized, middlewares.Project, update)
	}
}
