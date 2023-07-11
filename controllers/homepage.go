package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type JSearchResponse struct {
	Status     string `json:"status"`
	RequestID  string `json:"request_id"`
	Parameters struct {
		Query    string `json:"query"`
		Page     int    `json:"page"`
		NumPages int    `json:"num_pages"`
	} `json:"parameters"`
	Data []struct {
		EmployerName         string  `json:"employer_name"`
		EmployerLogo         string  `json:"employer_logo"`
		EmployerWebsite      string  `json:"employer_website"`
		EmployerCompanyType  string  `json:"employer_company_type"`
		JobPublisher         string  `json:"job_publisher"`
		JobID                string  `json:"job_id"`
		JobEmploymentType    string  `json:"job_employment_type"`
		JobTitle             string  `json:"job_title"`
		JobApplyLink         string  `json:"job_apply_link"`
		JobApplyIsDirect     bool    `json:"job_apply_is_direct"`
		JobApplyQualityScore float64 `json:"job_apply_quality_score"`
		JobDescription       string  `json:"job_description"`
	} `json:"data"`
}

// GetHomePage handles the root endpoint "/"
func GetHomePage(c *gin.Context) {
	url := "https://jsearch.p.rapidapi.com/search?query=Python%20developer%20in%20Texas%2C%20USA&page=1&num_pages=1"

	req, _ := http.NewRequest("GET", url, nil)

	apiKey := os.Getenv("RAPIDAPI_KEY") // Read the API key from the environment variable

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "jsearch.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error making the request",
		})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error reading the response body",
		})
		return
	}

	var response JSearchResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error parsing the response",
		})
		return
	}

	fmt.Printf("%+v\n", response)

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Gemstack community!",
		"data":    response.Data,
	})
}
