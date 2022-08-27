package main

import (
	"regentemr/controllers"
	"regentemr/database"
	"regentemr/middlewares"
	"regentemr/models"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
		}

		protected := api.Group("/auth").Use(middlewares.Auth())
		{
			protected.GET("/profile", controllers.Profile)
		}
	}

	return r
}

func main() {

	dsn := database.DB_USERNAME + ":" + database.DB_PASSWORD + "@tcp" + "(" + database.DB_HOST + ":" + database.DB_PORT + ")/" + database.DB_NAME + "?" + "parseTime=true&loc=Local"
	database.DBCon, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	database.DBCon.AutoMigrate(&models.User{})

	// // CREATE TEST DATA

	// // User Data
	database.DBCon.Create(&models.User{Username: "OBrien", Firstname: "Brian", Lastname: "Otieno", Phonenumber: "+254723328969", Email: "gebryo@intelligencia.com", Password: "$2a$14$SaSgFyNhW9ncAmMf19BTg.wSlAV2dctl/MXNsSdpYKupJE6AWAhpy", Role: "SADMIN", Approved: true})
	database.DBCon.Create(&models.User{Username: "jdoe", Firstname: "Jane", Lastname: "Doe", Phonenumber: "+254713846445", Email: "jane.doe.com", Password: "$2a$14$hTW2RgNLwKpnG4tNVKMZLed2thwKqtSuB7OOJhsJmKg8HD8/FRjkS", Role: "ADMIN", Approved: true})

	r := setupRouter()
	r.Run(":8085")
}
