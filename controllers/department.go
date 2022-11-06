package controllers

import (
	"log"
	"regentemr/models"

	"github.com/gin-gonic/gin"
)

func CreateDepartment(c *gin.Context) {
	var department models.Department

	err := c.ShouldBindJSON(&department)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{
			"msg": "invalid json input",
		})
		c.Abort()

		return
	}

	err = department.CreateDepartment()
	if err != nil {

		c.JSON(409, gin.H{
			"msg": "Department Exists",
		})
		c.Abort()

		return
	}

	c.JSON(200, department)

}
