package models

import (
	"github.com/jinzhu/gorm"
	"ngtium/libraries/common"
)

// Comment data model
type Comment struct {
	gorm.Model
	Message string `sql:"type:text;"`
	User    User   `gorm:"foreignkey:UserID"`
	UserID  uint
	Task    Task `json:"-";gorm:"foreignkey:TaskID"`
	TaskID  uint
}

// Serialize serializes comment data
func (p Comment) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"user_id":    p.User.Serialize(),
		"task_id":    p.TaskID,
		"message":    p.Message,
		"created_at": p.CreatedAt,
	}
}
