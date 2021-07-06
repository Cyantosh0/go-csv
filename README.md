# go-csv

### Easy CSV binding in Go!

Example that demonstrates reading and writing CSV files using https://github.com/gocarina/gocsv.

This Repo has three handlers:

**1. GET: /export-csv**
When a request is made to the API, the sample records of employee from sample_data.go file will be written over CSV file and the file will get saved in downloads folder.

**2. GET: /export-records**
Requesting to this API will provide the response as Blob that contains the sample records of employee from sample_data.go file.

**3. POST: /upload-csv**
Upload a csv file using formdata as body. Each records will get unmarshalled into Employee data struct.
