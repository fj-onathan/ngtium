package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Project{},
		&Task{},
		&Category{},
		&Comment{},
	)
	// set up foreign keys
	db.Model(&Project{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Task{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Model(&Task{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	db.Model(&Category{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Model(&Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Comment{}).AddForeignKey("task_id", "tasks(id)", "CASCADE", "CASCADE")
	fmt.Println("Auto Migration has beed processed")
}
