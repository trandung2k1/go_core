package switchs

import "fmt"

func Switchs() {
	// Program to print the day of the week using  switch case
	dayOfWeek := 3
	switch dayOfWeek {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
		fallthrough
	case 3:
		fmt.Println("Tuesday")
		fallthrough
	case 4:
		fmt.Println("Wednesday")
		fallthrough
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}

	// Result: Tuesday, Wednesday, Thursday
}
