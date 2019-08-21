package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"ngtium/database/models"
)

type User = models.User

func Project(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("project")
	user := c.MustGet("user").(User)

	var project models.Project
	if err := db.Where("id = ?", id).First(&project).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"success": false,
				"error":   "Project " + id + " not exists.",
			},
		)
		return
	}

	if project.UserID != user.ID {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"success": false,
				"error":   "Not allowed to access to project: " + id,
			},
		)
		return
	}
}

func Task(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("task")
	user := c.MustGet("user").(User)

	var task models.Task
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"success": false,
				"error":   "Task " + id + " not exists.",
			},
		)
		return
	}

	if task.UserID != user.ID {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"success": false,
				"error":   "Not allowed to access to task: " + id,
			},
		)
		return
	}
}

func Category(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("category")

	var category models.Category
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&category).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"success": false,
				"error":   "Category not exists.",
			},
		)
		return
	}
}

func Comment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(User)
	id := c.Param("comment")
	var comment models.Comment

	if err := db.Preload("User").Where("id = ?", id).First(&comment).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"success": false,
				"error":   "Comment not exists.",
			},
		)
		return
	}

	if comment.UserID != user.ID {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"success": false,
				"error":   "Not authorized to access that comment.",
			},
		)
		return
	}
}
