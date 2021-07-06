package csv

type Employee struct {
	ID   string `csv:"id"`
	Name string `csv:"name"`
	Age  int    `csv:"age"`
}

func getData() []Employee {
	return []Employee{
		{
			ID:   "E01",
			Name: "Daniel",
			Age:  27,
		},
		{
			ID:   "E02",
			Name: "Robin",
			Age:  25,
		},
		{
			ID:   "E03",
			Name: "Emmi",
			Age:  26,
		},
		{
			ID:   "E024",
			Name: "Samuel",
			Age:  25,
		},
	}
}
