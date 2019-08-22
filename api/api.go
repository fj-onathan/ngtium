package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ngtium/api/v1"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	// index documentation
	r.LoadHTMLGlob("html/*/*")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "api")
	})
	r.Static("/assets", "./html/assets")
	r.GET("api", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":       "ngtium documentation",
			"description": "GO API with JWT auth & mySQL",
			"github": gin.H{
				"name": "Github",
				"link": "https://github.com/dev-fjonathan/ngtium",
			},
			"documentation": gin.H{
				"name": "Documentation",
				"link": "https://github.com/dev-fjonathan/ngtium/tree/master/documentation",
			},
		})
	})
	// routing by grouping
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api)
	}
}
