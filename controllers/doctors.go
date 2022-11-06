package controllers

import (
	"log"
	"regentemr/models"

	"github.com/gin-gonic/gin"
)

func NewDoctor(c *gin.Context) {
	var doctor models.Doctor

	err := c.ShouldBindJSON(&doctor)
	if err != nil {
		log.Println(err)

	}
	c.JSON(400, gin.H{
		"msg": "invalid input",
	})
	c.Abort()
}
