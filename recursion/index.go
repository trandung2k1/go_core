package recursion

import "fmt"

func sum(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + sum(n-1)
	}

}
func display(n int) {
	if n > 0 {
		fmt.Println(n)
		display(n - 1)
	} else {
		fmt.Println("Stop")
	}
}

func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * sum(n-1)
	}
}
func Recursion() {
	fmt.Println(sum(3)) //6
	display(10)
	fmt.Println(factorial(4)) //24
}
