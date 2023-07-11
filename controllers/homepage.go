package controllers

import (
	"github.com/gin-gonic/gin"
	g "github.com/serpapi/google-search-results-golang"
    "fmt"
    "os"
    "github.com/joho/godotenv"
    "log"
)

// ROUTE: /
func GetHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Gemstack community!",
	})
	}

func goDotEnvVariable(key string) string {

      // load .env file
      err := godotenv.Load(".env")

      if err != nil {
        log.Fatalf("Error loading .env file")
      }

      return os.Getenv(key)
      }

func GoogleSearchAPI(c *gin.Context){

  searchTerm := "Barista"

  parameter := map[string]string{
    "engine": "google_jobs",
    "google_domain": "google.com",
    "q": searchTerm,
  }

  api_key := goDotEnvVariable("secret_api_key")

  search := g.NewGoogleSearch(parameter, api_key)
  results, err := search.GetJSON()
      if err != nil {
           fmt.Println(err.Error())
      }

      c.JSON(200, gin.H{
            "message": results,
        })
}
