package controllers

import (
	"log"
	"regentemr/auth"
	"regentemr/database"
	"regentemr/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Signup creates a user in db
func Signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()

		return
	}

	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())

		c.JSON(500, gin.H{
			"msg": "error hashing password",
		})
		c.Abort()

		return
	}

	err = user.CreateUserRecord()
	if err != nil {
		log.Println(err)

		c.JSON(409, gin.H{
			"msg": "User Exist",
		})
		c.Abort()

		return
	}

	c.JSON(200, user)
}

// LoginPayload login body
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Firstname   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Username    string `json:"username"`
	Token       string `json:"token"`
	Role        string `json:"role"`
	Phonenumber string `json:"phonenumber"`
}

// Login logs users in
func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result := database.DBCon.Where("username = ?", payload.Username).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "AA6161EE632FAB950F1FB6A79F946ABBEF759265",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Username)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Firstname:   user.Firstname,
		LastName:    user.Lastname,
		Username:    user.Username,
		Token:       signedToken,
		Role:        user.Role,
		Phonenumber: user.Phonenumber,
	}

	c.JSON(200, tokenResponse)

	return
}

// returns user data
func Profile(c *gin.Context) {
	var user models.User

	username, _ := c.Get("username") // from the authorization middleware

	result := database.DBCon.Where("username = ?", username.(string)).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get user profile",
		})
		c.Abort()
		return
	}

	user.Password = ""

	c.JSON(200, user)

	return
}
