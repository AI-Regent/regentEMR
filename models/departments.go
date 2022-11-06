package models

import (
	"regentemr/database"

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
}

// creates a department
func (department *Department) CreateDepartment() error {
	result := database.DBCon.Create(&department)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
