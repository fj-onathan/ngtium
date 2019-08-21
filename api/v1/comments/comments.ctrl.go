package comments

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"ngtium/database/models"
	"ngtium/libraries/common"
	"strconv"
)

// Comment type alias
type Comment = models.Comment

// User type alias
type User = models.User

// JSON type alias
type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")
	task := c.Param("task")

	var comments []Comment

	if cursor == "" {
		if err := db.Preload("User").Limit(10).Order("id desc").Where("task_id = ?", task).Find(&comments).Error; err != nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{
					"success": false,
					"error":   "Not founded recent comments in task.",
				},
			)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Preload("User").Limit(10).Order("id desc").Where(condition, cursor).Where("task_id = ?", task).Find(&comments).Error; err != nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{
					"success": false,
					"error":   "Not founded recent comments in task.",
				},
			)
			return
		}
	}

	length := len(comments)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = comments[i].Serialize()
	}

	c.JSON(http.StatusOK, common.JSON{
		"success": true,
		"user":    serialized,
	})
}

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	task := c.Param("task")
	type RequestBody struct {
		Message string `json:"message" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"success": false,
			"error":   "Required inputs cannot be empty",
		})
		return
	}

	tID, _ := strconv.ParseUint(task, 10, 0)
	user := c.MustGet("user").(User)
	comment := Comment{
		Message: requestBody.Message,
		User:    user,
		TaskID:  uint(tID),
	}
	db.NewRecord(comment)
	db.Create(&comment)
	c.JSON(http.StatusOK, common.JSON{
		"data":    comment.Serialize(),
		"success": true,
	})
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("comment")

	type RequestBody struct {
		Message string `json:"message" binding:"required"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"success": false,
				"error":   "Required inputs cannot be empty.",
			},
		)
		return
	}

	var comment Comment
	db.Preload("User").Where("id = ?", id).First(&comment)

	comment.Message = requestBody.Message
	db.Save(&comment)
	c.JSON(http.StatusOK, common.JSON{
		"success": true,
		"data":    comment.Serialize(),
	})
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("comment")
	var comment Comment
	db.Where("id = ?", id).First(&comment)
	db.Delete(&comment)
	c.JSON(http.StatusOK, common.JSON{
		"success": true,
	})
}
