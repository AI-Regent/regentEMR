package database

import (
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "P@55w0rd"
const DB_NAME = "regentemr"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

var (
	// DBCon is the connection handle
	// for the database
	DBCon *gorm.DB
)

// var DB *gorm.DB

// var dsn = DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

// func ConectDatabase() {
// 	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
// 	if err != nil {

// 		panic("Failed To Open Database!")
// 	}
// 	// database.AutoMigrate(&models.Department{})
// 	// database.AutoMigrate(&models.User{})
// 	// database.AutoMigrate(&models.Doctor{})
// 	// database.AutoMigrate(&models.Department{})

// 	// database.Create(&models.User{Username: "OBrien", Firstname: "Brian", Lastname: "Otieno", Phonenumber: "+254723328969", Email: "gebryo@intelligencia.com", Password: "$2a$14$SaSgFyNhW9ncAmMf19BTg.wSlAV2dctl/MXNsSdpYKupJE6AWAhpy", Role: "SADMIN", Approved: true})
// 	// database.Create(&models.User{Username: "jdoe", Firstname: "Jane", Lastname: "Doe", Phonenumber: "+254713846445", Email: "jane.doe.com", Password: "$2a$14$hTW2RgNLwKpnG4tNVKMZLed2thwKqtSuB7OOJhsJmKg8HD8/FRjkS", Role: "ADMIN", Approved: true})

// 	DB = database
// }
