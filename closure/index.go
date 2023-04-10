package closure

import "fmt"

//return func
func msg() func() {
	return func() {
		fmt.Println("Hello Dung")
	}
}

func displayNumbers() func() int {
	number := 0
	return func() int {
		number++
		return number
	}

}
func Closure() {
	msg() //Hello Dung
	num1 := displayNumbers()
	//context1
	fmt.Println(num1()) //1
	fmt.Println(num1()) //2
	fmt.Println(num1()) //3

	//context2
	num2 := displayNumbers()
	fmt.Println(num2()) //1
	fmt.Println(num2()) //2
}
