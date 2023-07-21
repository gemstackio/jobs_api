package main

import (
	DbConnection "github.com/gemstack/jobs-api/dbconnection"
	"github.com/gemstack/jobs-api/routes"
)

func main() {

	DbConnection.DbConnection()

	//r := gin.Default()

	// Define your routes and handlers here
	r := routes.SetupRoutes()

	r.Run(":8080")

}
