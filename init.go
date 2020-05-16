package main
import (
	"puregoboilerbplate/datastructure"
)

func initStruct() []datastructure.Employee{
	karyawan := []datastructure.Employee{
		datastructure.Employee{
			ID: "01",
			Name: "Anas",
			Role: "Backend",
		},
		datastructure.Employee{
			ID: "02",
			Name: "Dhiko",
			Role: "Backend",
		},
		datastructure.Employee{
			ID: "03",
			Name: "Asdi",
			Role: "Frontend",
		},
	}
	return karyawan
}