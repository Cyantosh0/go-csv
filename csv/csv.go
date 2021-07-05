package csv

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func ExportInCSV(c *gin.Context) {
	file, err := ioutil.TempFile("downloads", "employee-*.csv")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	data := getData()
	gocsv.MarshalFile(&data, file)

	c.JSON(200, gin.H{"message": "success"})
}

func ExportRecords(c *gin.Context) {
	data := getData()

	csvBytes, err := gocsv.MarshalBytes(&data)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.WriteHeader(200)
	c.Writer.Write(csvBytes)
}
