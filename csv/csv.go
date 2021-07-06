package csv

import (
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

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

func UploadCSV(c *gin.Context) {
	type csvUploadInput struct {
		CsvFile *multipart.FileHeader `form:"file" binding:"required"`
	}

	var input csvUploadInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if filepath.Ext(input.CsvFile.Filename) != ".csv" && input.CsvFile.Header.Get("Content-Type") != "text/csv" {
		c.JSON(400, gin.H{"error": "upload a csv file"})
		return
	}

	f, err := input.CsvFile.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var employee []Employee
	gocsv.UnmarshalBytes(fileBytes, &employee)

	c.JSON(200, employee)
}
