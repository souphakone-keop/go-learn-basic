package slices

import "fmt"

func Slices() {
	mySlice := []int{1, 2, 3, 4}

	indexTwo := mySlice[1:3]
	fmt.Println("slices :", indexTwo)

}
