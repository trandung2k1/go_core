package packages

import (
	"fmt"
	"golang/packages/calculator"
	"math"
	"strings"
)

func Package() {
	// package:
	// math: Sqrt(), Cbrt(), Max(), Min(), Mod()
	// string: Compare(),Contains(), Count(), Join(), ToLower(), ToLower()
	// fmt: Print(), Scan(), Scanf(), Scanln(), Printf(), Println()

	fmt.Println("Hello world")
	var x int
	fmt.Printf("Enter x: ")
	fmt.Scan(&x)
	fmt.Println(x)
	fmt.Println(math.Sqrt(float64(x)))

	//use slice
	var str = []string{"Hello", "world"}
	fmt.Println(strings.Join(str, " ")) //Hello world
	fmt.Printf("%T", str)

	//use declare package

	number1 := 9
	number2 := 5
	fmt.Println()
	// use the add function of calculator package
	fmt.Println(calculator.Add(number1, number2)) //14
	// use the subtract function of calculator package
	fmt.Println(calculator.Subtract(number1, number2)) //4
}
