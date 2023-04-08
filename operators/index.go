package operators

import "fmt"

func Operators() {
	// Arithmetic  Operator: +, -, *, /, %
	num1 := 6
	num2 := 2
	sum := num1 + num2

	fmt.Printf("%d + %d = %d\n", num1, num2, sum) // 6 + 2 = 8
	difference := num1 - num2
	fmt.Printf("%d - %d = %d\n", num1, num2, difference) // 6 - 2 = 4
	product := num1 * num2
	fmt.Printf("%d * %d = %d\n", num1, num2, product) // 6 * 2 = 12

	num3 := 11
	num4 := 4
	quotient := num3 / num4
	fmt.Printf("%d / %d = %d\n", num3, num4, quotient) //11 / 4 = 2

	//fix issue
	num5 := 11.0
	num6 := 4.0
	fmt.Println(num5 / num6) // 2.75

	var num7 float32 = 11
	var num8 float32 = 4
	fmt.Println(num7 / num8) //2.75

	//Modulo

	fmt.Println(11 % 4) //3
	//Golang not prevfix operator and **
	//Rust, swift 3 not prevfix operator

	//Assignment operator
	num := 6
	res := num
	fmt.Println(res) //6

	//Compound assignment operators
	num += 10
	//Relational Operators in Golang: ==, !=, >, <, >=, <= (boolean)


	//Logical operators: &&, || , !
}
