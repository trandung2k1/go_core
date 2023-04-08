package ifelses

import "fmt"

func IfElse() {
	number := -115
	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	} else {
		fmt.Printf("%d is a negative number\n", number)
	}
	fmt.Println("Out the loop")

	var number1 int = 1
	var number2 int = 2
	if number1 == number2 {
		fmt.Printf("Result: %d == %d", number1, number2)

	} else if number1 > number2 {
		fmt.Printf("Result: %d > %d", number1, number2)

	} else {
		fmt.Printf("Result: %d < %d", number1, number2)

	}

	//else if
	//Nested if,..
}
