package booleanexpression

import "fmt"

func BooleanExpression() {
	var boolTrue bool = true
	var boolFalse bool = false
	fmt.Println(boolFalse, boolTrue)

	number1 := 9
	number2 := 3
	result := number1 > number2

	fmt.Println(result) // true

	//Relational Operators in Golang: ==, !=, >, <, >=, <=
	// Logical Operators in Golang: &&, ||, !
}
