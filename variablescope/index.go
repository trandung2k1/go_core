package variablescope

import "fmt"

//global

var x = 10

func addNumbers() {
	//local
	var sum int
	sum = 5 + 9
	fmt.Println(sum) //14
	fmt.Println(x)   //10
}

func VariableScope() {
	addNumbers()
	fmt.Println(x) //10
}
