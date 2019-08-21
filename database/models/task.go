package models

import (
	"github.com/jinzhu/gorm"
	"ngtium/libraries/common"
)

// Task data model
type Task struct {
	gorm.Model
	Title       string   `sql:"type:varchar(100);"`
	Description string   `sql:"type:text;"`
	Priority    string   `sql:"type:varchar(20);"`
	Status      string   `sql:"type:varchar(20);"`
	CategoryID  uint     `sql:"default: null;"`
	Category    Category `json:"-"`
	User        User     `gorm:"foreignkey:UserID"`
	UserID      uint
	Project     Project `gorm:"foreignkey:ProjectID"`
	ProjectID   uint
}

// Project fields
type HProject struct {
	ID    uint
	Title string
}

// Category fields
type HCategory struct {
	ID          uint
	Name        string
	Description string
}

// Serialize serializes task data
func (p *Task) Serialize() common.JSON {
	return common.JSON{
		"id":          p.ID,
		"user_id":     p.User.Serialize(),
		"project":     HProject{ID: p.Project.ID, Title: p.Project.Title},
		"category":    HCategory{ID: p.CategoryID, Name: p.Category.Name, Description: p.Category.Description},
		"title":       p.Title,
		"description": p.Description,
		"priority":    p.Priority,
		"status":      p.Status,
		"created_at":  p.CreatedAt,
	}
}

func (p *Task) Read() common.JSON {
	return common.JSON{
		"id":          p.ID,
		"user":        p.User.Serialize(),
		"project_id":  p.Project.ID,
		"category_id": p.CategoryID,
		"title":       p.Title,
		"description": p.Description,
		"priority":    p.Priority,
		"status":      p.Status,
		"created_at":  p.CreatedAt,
	}
}
