package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"ngtium/database/models"
	"ngtium/libraries/common"
	"strconv"
)

// Project type alias
type Project = models.Project
type Task = models.Task

// Owner type alias
type Owner = models.User

// JSON type alias
type JSON = common.JSON

func listRecent(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var tasks []Task

	if cursor == "" {
		if err := db.Preload("Project").Preload("User").Limit(10).Order("id desc").Find(&tasks).Error; err != nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{
					"success": false,
					"error":   "Not founded recent tasks in project.",
				},
			)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Preload("Project").Preload("User").Limit(10).Order("id desc").Where(condition, cursor).Find(&tasks).Error; err != nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{
					"success": false,
					"error":   "Not founded recent tasks in project.",
				},
			)
			return
		}
	}

	length := len(tasks)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = tasks[i].Serialize()
	}

	c.JSON(http.StatusOK, common.JSON{
		"data":    serialized,
		"success": true,
	})
}

func list(c *gin.Context) {

	id := c.Param("task")
	action := c.Param("action")

	switch action {

	case "list":
		listByProject(c, id)

	case "read":
		read(c, id)

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

func read(context *gin.Context, id string) {
	db := context.MustGet("db").(*gorm.DB)
	var task Task
	db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&task)
	context.JSON(http.StatusOK, common.JSON{
		"data":    task.Serialize(),
		"success": true,
	})
}

func listByProject(context *gin.Context, id string) {
	db := context.MustGet("db").(*gorm.DB)
	var tasks []Task

	db.Set("gorm:auto_preload", true).Where("project_id = ?", id).Find(&tasks)

	length := len(tasks)
	serialized := make([]JSON, length, length)
	for i := 0; i < length; i++ {
		serialized[i] = tasks[i].Read()
	}

	context.JSON(http.StatusOK, common.JSON{
		"data":    serialized,
		"success": true,
	})
}

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("project")
	user := c.MustGet("user").(Owner)

	var project Project
	db.Where("id = ?", id).First(&project)

	type RequestBody struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Priority    string `json:"priority"`
		CategoryID  uint   `json:"category_id"`
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

	pID, _ := strconv.ParseUint(id, 10, 0)
	task := Task{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Priority:    requestBody.Priority,
		CategoryID:  requestBody.CategoryID,
		ProjectID:   uint(pID),
		User:        user,
	}
	db.NewRecord(task)
	db.Create(&task)
	c.JSON(http.StatusOK, common.JSON{
		"data":    task.Serialize(),
		"success": true,
	})
}

func update(c *gin.Context) {
	id := c.Param("task")
	action := c.Param("action")
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Status      string `json:"status" binding:"required"`
	}
	var requestBody RequestBody
	var task Task

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

	db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&task)

	switch action {

	case "update":
		task.Title = requestBody.Title
		task.Description = requestBody.Description
	case "status":
		task.Status = requestBody.Status
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

	db.Save(&task)
	c.JSON(200, common.JSON{
		"data":    task.Serialize(),
		"success": true,
	})
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("task")
	var task Task
	db.Where("id = ?", id).First(&task)
	db.Delete(&task)
	c.JSON(http.StatusOK, common.JSON{
		"success": true,
	})
}
