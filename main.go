package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"cyantosh/go-csv/csv"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	}))

	router.GET("/export-csv", csv.ExportInCSV)

	router.GET("/export-records", csv.ExportRecords)

	router.Run(":5000")
}
