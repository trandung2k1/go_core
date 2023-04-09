package functions

import "fmt"

func greet() {
	fmt.Println("Good morning")
}

func addNumbers(a int, b int) int {
	return a + b
}

//Return multiple value
func caculate(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	difference := n1 - n2
	return sum, difference
}

//Code Reusability

func getSquare(n int) int {
	return n * n
}
func Functions() {
	greet()
	sum := addNumbers(1, 2)
	fmt.Println(sum) //3
	s, d := caculate(2, 3)
	fmt.Println(s)
	fmt.Println(d)

	fmt.Println(getSquare(2))
	fmt.Println(getSquare(3))
}
