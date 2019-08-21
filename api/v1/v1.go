package apiv1

import (
	"github.com/gin-gonic/gin"
	"ngtium/api/v1/auth"
	"ngtium/api/v1/categories"
	"ngtium/api/v1/comments"
	"ngtium/api/v1/projects"
	"ngtium/api/v1/tasks"
	"ngtium/database"
	"ngtium/database/models"
)

// Ping connection
func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "do a pong",
	})
}

// Run migration
func migration(c *gin.Context) {
	// initializes database
	db, err := database.Initialize()
	if err != nil {
		panic(err)
	}
	models.Migrate(db)
	c.JSON(200, gin.H{
		"message": "Migration of database, done.",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		// Helpers routing
		v1.GET("/ping", ping)
		v1.GET("/migration", migration)

		// API routing
		auth.ApplyRoutes(v1)
		projects.ApplyRoutes(v1)
		tasks.ApplyRoutes(v1)
		categories.ApplyRoutes(v1)
		comments.ApplyRoutes(v1)
	}
}
