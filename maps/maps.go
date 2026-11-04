package maps

import "fmt"

func Maps() {
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 8
	myMap["orange"] = 10
	delete(myMap, "zd")

	val, ok := myMap["pear"]
	if ok {
		fmt.Println("pear is value:", val)
	} else {
		fmt.Println("pear not fount inmap")
	}
	/*
		type Student struct {
		   		Name   string
		   		Weight int
		   	}
		   	students := make(map[string]Student)
		   	students["st01"] = Student{Name: "souphakone", Weight: 70}
		   	students["st02"] = Student{Name: "surinm", Weight: 45}
		   	students["st03"] = Student{Name: "spaycial", Weight: 64}
		   	fmt.Println("studentt map: ", students)
	*/
	type Student struct {
		name   string
		weight int
	}
	students := make(map[string]Student)
	students["st1"] = Student{name: "soup", weight: 64}
	students["st2"] = Student{name: "boss", weight: 54}
	students["st3"] = Student{name: "ssess", weight: 74}
	fmt.Println("syuden map : ", students)
}
