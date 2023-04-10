package anonymousfunction

import "fmt"

func findSquare(num int) int {
	return num * num
}
func Function() {
	var greet = func() {
		fmt.Println("Hello")
	}
	greet() //Hello
	sum := func(n1, n2 int) {
		s := n1 + n2
		fmt.Println("Sum is:", s)
	}
	sum(1, 2) //3
	area := func(length, breadth int) int {
		return length * breadth
	}
	fmt.Println("The area of rectangle is", area(3, 4)) //12

	result := findSquare(area(1, 2))
	fmt.Println("Result is:", result) //4

}
