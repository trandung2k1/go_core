package deferpanicrecover

import "fmt"

func division(num1, num2 int) {
	// if num2 is 0, program is terminated due to panic
	defer handlePanic()
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}
func handlePanic() {
	a := recover()
	if a != nil {
		fmt.Println("RECOVER", a)
	}
}
func Main() {
	// defer fmt.Println("Hi")
	// fmt.Println("Hello")
	//rs Hello -> Hi

	//Example: LIFO (Last In First Out)
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	//Three Two One

	// End the execution

	division(4, 2) // Result:  2
	division(8, 0) // end execution
	division(2, 8)

	//recover  :prevents the termination
}
