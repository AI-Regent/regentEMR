package models

import (
	"regentemr/database"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User defines the user in db
type User struct {
	gorm.Model
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Username    string    `json:"username" gorm:"unique"`
	Phonenumber string    `json:"phonenumber"`
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"password"`
	Role        string    `json:"role" validate:"required, eq=ADMIN|eq=QA|eq=QC|eq=DA|eq=MONITOR|eq=USER|eq=SADMIN"`
	Approved    bool      `json:"approved"`
	Created_at  time.Time `json:"created_at" gorm:"column:created_at; type:timestamp; default: NOW(); not null; <-:create"`
	Updated_at  time.Time `json:"updated_at" gorm:"column:created_at; type:timestamp; default: NOW(); <-:update"`
}

// creates a user
func (user *User) CreateUserRecord() error {
	result := database.DBCon.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
