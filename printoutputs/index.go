package printoutputs

import "fmt"

func PrintOutput() {

	//fmt.Print(string, variables, mutiple variables)
	name, age := "Dung", 21
	fmt.Print("Hi")      //Hi
	fmt.Print(name, age) // Dung 21

	//fmt.Println():  prints a new line at the end by default, multiple variables create new space
	fmt.Println("Hi:", name) //Hi: Dung

	//fmt.Printf()

	currentAge := 21
	fmt.Printf("%d\n", currentAge) //
	//Note: int - %d, float - %g, string - %s, bool - %t, type - %T

	var salary float32 = 20.5
	fmt.Printf("%g\n", salary) //20.5
	fmt.Printf("%f\n", salary) //20.500000

	var isLoading bool = false
	fmt.Printf("%t", isLoading) // false

	//use print() replace fmt.Print(), println() replace fmt.Println()
	// use print() and println() debug
}
