// Program to check if the day is a weekend or a weekday

package switchs

import "fmt"

func Example() {
	dayOfWeek := "Sunday"

	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	default:
		fmt.Println("Invalid day")
	}

	//Result: Weekend
}
