package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gemstack/jobs-api/controllers"
	"github.com/gemstack/jobs-api/utils"
)

func main() {
	r := gin.Default()

	// Enable CORS middleware
	r.Use(utils.CORSMiddleware())

	// Define your routes and handlers here
	r.GET("/", controllers.GetHomePage)

	r.Run(":8080")
}
