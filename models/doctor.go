package models

import (
	"regentemr/database"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Department  string `json:"department"`
	Description string `json:"description"`
}

func (doctor *Doctor) CreateDoctor() error {
	if err := database.DBCon.Create(&doctor).Error; err != nil {
		return err
	}
	return nil
}
