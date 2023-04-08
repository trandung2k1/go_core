package takeinputs

import "fmt"

func TakeInputs() {
	// var name string
	// var age int
	// fmt.Print("Enter your name: ")
	// fmt.Scan(&name)
	// fmt.Println("Enter your name and age: ")

	// Enter 1 line or multiple lines
	// fmt.Scan(&name, &age)

	//Enter data 1 line
	// fmt.Scanln(&name, &age)

	//Enter data 1 line
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Println(name, age)

	//float and boolean

	var isLoading bool
	fmt.Println("Enter loading: ")
	fmt.Scanf("%t", &isLoading)
	fmt.Println(isLoading)

	var salary float64
	fmt.Println("Enter salary: ")
	fmt.Scanf("%g", &salary)
	fmt.Println(salary)
}
