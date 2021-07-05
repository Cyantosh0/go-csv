package main

import (
	"github.com/gin-gonic/gin"

	"cyantosh/go-csv/csv"
)

func main() {
	router := gin.Default()

	router.GET("/export-csv", csv.ExportInCSV)

	router.Run(":5000")
}
