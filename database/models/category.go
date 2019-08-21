package models

import (
	"github.com/jinzhu/gorm"
	"ngtium/libraries/common"
)

type Counters struct {
	All  int
	Done int
}

// Category data model
type Category struct {
	gorm.Model
	Name        string  `sql:"type:varchar(100);"`
	Description string  `sql:"type:text;"`
	Project     Project `gorm:"foreignkey:ProjectID"`
	ProjectID   uint
	Counters    Counters `sql:"-"`
}

// Serialize serializes category data
func (p Category) Serialize() common.JSON {
	return common.JSON{
		"id":          p.ID,
		"project_id":  p.ProjectID,
		"name":        p.Name,
		"description": p.Description,
		"created_at":  p.CreatedAt,
		"counters":    p.Counters,
	}
}
