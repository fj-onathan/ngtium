package comments

import (
	"github.com/gin-gonic/gin"
	"ngtium/libraries/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	comments := r.Group("/comments")
	{
		comments.POST("/:task", middlewares.Authorized, middlewares.Task, create)
		comments.GET("/:task", list)
		comments.DELETE("/:comment", middlewares.Authorized, middlewares.Comment, remove)
		comments.PATCH("/:comment", middlewares.Authorized, middlewares.Comment, update)
	}
}
