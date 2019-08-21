package api

import (
	"github.com/gin-gonic/gin"
	"ngtium/api/v1"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api)
	}
}
