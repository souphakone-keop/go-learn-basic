package keopheth

import "fmt"

func Keopheth() {
	var dayOfweek int = 3
	switch dayOfweek {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("allDay")

	}

	for i := 1; i <= dayOfweek; i++ {
		fmt.Println("number", i)
	}

	j := 1
	for {
		fmt.Println("j:", j)
		j++
		if j > 3 {
			break
		}
	}
}
