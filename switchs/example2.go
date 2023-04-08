package switchs

import "fmt"

func Example2() {
	numberOfDay := 28
	//default switch true
	switch {
	case 28 == numberOfDay:
		fmt.Println("It's February")
	default:
		fmt.Println("Not January")
	}
}
