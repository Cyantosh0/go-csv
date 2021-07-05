package csv

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func ExportInCSV(c *gin.Context) {
	file, err := ioutil.TempFile("downloads", "employee-*.csv")
	if err != nil {
		fmt.Println("Error occurred : ", err)
	}
	defer file.Close()

	data := getData()
	gocsv.MarshalFile(&data, file)

	c.JSON(200, gin.H{"message": "success"})
}
