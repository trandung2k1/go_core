package variables

import "fmt"

func Variable() {
	var number int = 10
	var num = 10
	n := 10
	fmt.Printf("\nNumber =: %d", number) //10
	fmt.Printf("\nNum = %d", num)        //10
	fmt.Printf("\nn = %d", n)            //10
	var count int
	fmt.Printf("\ncount = %d", count) // 0
	// var count2 // error: missing variable type or initialization
	// change value
	n = 100
	fmt.Printf("\nn = %d", n) //100
	// n = "Hello" // error: cannot use "Hello" (untyped string constant) as int value in assignment
	//create multiple variables
	// var name, age = "Dung", 21
	name, age := "Dung", 21
	fmt.Printf("\nname = %s", name) //Dung
	fmt.Printf("\nage = %d", age)   //21
	//Type the variable
	fmt.Printf("\ntype name %T", name) //string

	//const
	const PI = 3.14
	// PI = 3.15 // error: cannot assign to PI
	fmt.Printf("\nPI = %f", PI)
	// const Speed := 10; // error: missing constant value
}
