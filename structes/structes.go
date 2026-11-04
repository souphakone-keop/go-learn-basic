package structes

import "fmt"

func Struct() {

	type Address struct {
		city    string
		country string
	}
	type Student struct {
		Name    string
		Height  float64
		Grade   string
		Address Address
	}

	var students [3]Student
	students[0] = Student{
		Name:   "Souphakone",
		Height: 183,
		Grade:  "A",
		Address: Address{
			city:    "pakse",
			country: "laos",
		},
	}
	students[1] = Student{
		Name:   "John",
		Height: 175,
		Grade:  "B",
	}
	for i := 0; i < len(students); i++ {
		fmt.Println("Student", i+1)
		fmt.Println("Name:", students[i].Name)
		fmt.Println("Height:", students[i].Height)
		fmt.Println("Grade:", students[i].Grade)
		fmt.Println("Grade:", students[i].Address)
	}

}
