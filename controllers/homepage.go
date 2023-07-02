package controllers

import (
	"github.com/gin-gonic/gin"
	g "github.com/serpapi/google-search-results-golang"
    "fmt"
)

// ROUTE: /
func GetHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Gemstack community!",
	})

	func GoogleSearchAPI(c *gin.Context){

      parameter := map[string]string{
        "engine": "google_jobs",
        "google_domain": "google.com",
        "q": "Barista",
        "api_key": "secret_api_key",
      }

      search := g.NewGoogleSearch(parameter, "secret_api_key")
      results, err := search.GetJSON()
          if err != nil {
               fmt.Println(err.Error())
          }

          c.JSON(200, gin.H{
          		"message": results,
          	})
}
