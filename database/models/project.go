package models

import (
	"github.com/jinzhu/gorm"
	"ngtium/libraries/common"
)

type Statistics struct {
	Tasks      int
	Percentage int
}

// Project data model
type Project struct {
	gorm.Model
	Title       string `sql:"type:varchar(50);"`
	Description string `sql:"type:text;"`
	User        User   `gorm:"foreignkey:UserID"`
	UserID      uint
	Statistics  Statistics `sql:"-"`
}

// Serialize serializes project data
func (p *Project) Serialize() common.JSON {
	return common.JSON{
		"id":          p.ID,
		"owner":       p.User.Serialize(),
		"title":       p.Title,
		"description": p.Description,
		"created_at":  p.CreatedAt,
		"stats":       p.Statistics,
	}
}
