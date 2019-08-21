package projects

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"ngtium/database/models"
	"ngtium/libraries/common"
)

// Project type alias
type Project = models.Project

// Task type alias
type Task = models.Task

// Owner type alias
type Owner = models.User

// JSON type alias
type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var projects []Project
	var tasks []Task

	if cursor == "" {
		if err := db.Preload("User").Limit(10).Order("id desc").Find(&projects).Error; err != nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{
					"success": false,
					"error":   "Not founded recent projects.",
				},
			)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Preload("User").Limit(10).Order("id desc").Where(condition, cursor).Find(&projects).Error; err != nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{
					"success": false,
					"error":   "Not founded recent projects.",
				},
			)
			return
		}
	}

	length := len(projects)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {

		var countTasks int
		var countTasksDone int
		var percentageTasks int

		db.Find(&Task{}).Where("project_id = ?", projects[i].ID).Find(&tasks)
		countTasks = len(tasks)
		stats := models.Statistics{}

		for t := 0; t < len(tasks); t++ {
			if tasks[t].Status == "done" {
				countTasksDone++
			}
		}

		if countTasks != 0 {
			percentageTasks = int((float64(countTasksDone) / float64(countTasks)) * 100)
		}

		stats.Tasks = countTasks
		stats.Percentage = percentageTasks

		projects[i].Statistics = stats
		serialized[i] = projects[i].Serialize()

	}

	c.JSON(http.StatusOK, serialized)
}

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "Required inputs cannot be empty",
		})
		return
	}

	user := c.MustGet("user").(Owner)
	project := Project{Title: requestBody.Title, Description: requestBody.Description, User: user}
	db.NewRecord(project)
	db.Create(&project)
	c.JSON(http.StatusOK,
		common.JSON{
			"success": true,
			"data":    project.Serialize(),
		},
	)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("project")

	var project Project
	db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&project)

	c.JSON(http.StatusOK, common.JSON{
		"success": true,
		"user":    project.Serialize(),
	})
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("project")

	type RequestBody struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"success": false,
				"error":   "Required inputs cannot be empty.",
			},
		)
		return
	}

	var project Project
	db.Preload("User").Where("id = ?", id).First(&project)

	project.Title = requestBody.Title
	project.Description = requestBody.Description
	db.Save(&project)
	c.JSON(http.StatusOK, common.JSON{
		"data":    project.Serialize(),
		"success": true,
	})
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("project")
	db.Where("id = ?", id).Delete(&Project{})
	db.Where("project_id = ?", id).Delete(&Task{})
	c.JSON(http.StatusOK, common.JSON{
		"success": true,
	})
}
