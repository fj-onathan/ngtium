package tasks

import (
	"github.com/gin-gonic/gin"
	"ngtium/libraries/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	tasks := r.Group("/tasks")
	{
		tasks.POST("/:project", middlewares.Authorized, middlewares.Project, create)
		tasks.GET("/", middlewares.Authorized, listRecent)
		tasks.GET("/:task/:action", middlewares.Authorized, middlewares.Task, list)
		tasks.DELETE("/:task", middlewares.Authorized, middlewares.Task, remove)
		tasks.PATCH("/:task/:action", middlewares.Authorized, middlewares.Task, update)
	}
}
