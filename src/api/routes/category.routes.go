package routes

import (
	"github.com/gin-gonic/gin"

	"grpcTreino/src/api/models"
)

func HomeRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the API",
	})
}

func GetAllCategories(c *gin.Context) {
	var categories models.Categories
	models.Connection().Find(&categories)
	c.JSON(200, gin.H{
		"categories": categories,
	})
}
