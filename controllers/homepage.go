package controllers

import (
	"github.com/gin-gonic/gin"
)

// ROUTE: /
func GetHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Gemstack community!",
	})
}
