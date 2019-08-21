package categories

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"ngtium/database/models"
	"ngtium/libraries/common"
	"strconv"
)

// Project type alias
type Category = models.Category
type Task = models.Task
type Project = models.Project

// Owner type alias
type Owner = models.User

// JSON type alias
type JSON = common.JSON

func list(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	action := c.Param("action")
	id := c.Param("category")

	switch action {

	case "list":

		var categories []Category
		var tasks []Task

		cursor := c.Query("cursor")
		recent := c.Query("recent")

		if cursor == "" {
			if err := db.Limit(10).Order("id desc").Where("project_id = ?", id).Find(&categories).Error; err != nil {
				c.AbortWithStatusJSON(
					http.StatusNotFound,
					gin.H{
						"success": false,
						"error":   "Not founded recent categories in project.",
					},
				)
				return
			}
		} else {
			condition := "id < ?"
			if recent == "1" {
				condition = "id > ?"
			}
			if err := db.Limit(10).Order("id desc").Where(condition, cursor).Find(&categories).Error; err != nil {
				c.AbortWithStatusJSON(
					http.StatusNotFound,
					gin.H{
						"success": false,
						"error":   "Not founded recent categories in project.",
					},
				)
				return
			}
		}

		length := len(categories)
		serialized := make([]JSON, length, length)

		for i := 0; i < length; i++ {

			var countTasks int
			var countTasksDone int

			db.Find(&Task{}).Where("category_id = ?", categories[i].ID).Find(&tasks)
			countTasks = len(tasks)
			counters := models.Counters{}

			for t := 0; t < len(tasks); t++ {
				if tasks[t].Status == "done" {
					countTasksDone++
				}
			}

			counters.All = countTasks
			counters.Done = countTasksDone

			categories[i].Counters = counters
			serialized[i] = categories[i].Serialize()

		}

		c.JSON(http.StatusOK, common.JSON{
			"success": true,
			"user":    serialized,
		})

	case "read":

		db := c.MustGet("db").(*gorm.DB)
		var category Category

		db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&category)
		c.JSON(http.StatusOK, common.JSON{
			"success": true,
			"user":    category.Serialize(),
		})

	default:
		c.AbortWithStatusJSON(
			http.StatusMethodNotAllowed,
			gin.H{
				"success": false,
				"error":   "Not allowed to use that action.",
			},
		)
		return
	}

}

func create(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("project")

	type RequestBody struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"success": false,
			"error":   "Required inputs cannot be empty",
		})
		return
	}
	pID, _ := strconv.ParseUint(id, 10, 0)
	category := Category{
		Name:        requestBody.Name,
		Description: requestBody.Description,
		ProjectID:   uint(pID),
	}
	db.NewRecord(category)
	db.Create(&category)
	c.JSON(http.StatusOK, common.JSON{
		"data":    category.Serialize(),
		"success": true,
	})
}

func update(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("category")
	project := c.Param("project")

	type RequestBody struct {
		Name        string `json:"name" binding:"required"`
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

	var category Category

	db.Where("id = ?", id).First(&category)

	projectID, _ := strconv.ParseUint(project, 10, 0)
	if uint(projectID) != category.ProjectID {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"success": false,
				"error":   "The category don't match with category",
			},
		)
		return
	}

	category.Name = requestBody.Name
	category.Description = requestBody.Description
	db.Save(&category)
	c.JSON(http.StatusOK, common.JSON{
		"success": true,
		"data":    category.Serialize(),
	})
}

func remove(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("category")
	project := c.Param("project")

	var category Category
	db.Where("id = ?", id).First(&category)

	projectID, _ := strconv.ParseUint(project, 10, 0)
	if uint(projectID) != category.ProjectID {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"success": false,
				"error":   "The category don't match with category",
			},
		)
		return
	}

	db.Delete(&category)
	db.Where("category_id = ?", id).Delete(&Task{})
	c.JSON(http.StatusOK, common.JSON{
		"success": true,
	})
}
