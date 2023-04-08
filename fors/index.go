package fors

import "fmt"

func Fors() {
	var sum int = 0

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum) //55
}
