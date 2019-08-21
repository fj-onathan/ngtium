package models

import (
	"github.com/jinzhu/gorm"
	"ngtium/libraries/common"
)

// User data model
type User struct {
	gorm.Model
	Username     string
	DisplayName  string
	PasswordHash string
	Email        string
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"username":     u.Username,
		"email":        u.Email,
		"display_name": u.DisplayName,
		"avatar":       string(u.DisplayName[0:2]),
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Username = m["username"].(string)
	u.DisplayName = m["display_name"].(string)
}
