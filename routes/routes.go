package routes

import (
	"github.com/gemstack/jobs-api/controllers"
	"github.com/gemstack/jobs-api/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Enable CORS middleware
	r.Use(utils.CORSMiddleware())
	r.GET("/gjobs", controllers.GoogleSearchAPI)
	r.GET("/", controllers.GetHomePage)

	return r
}
