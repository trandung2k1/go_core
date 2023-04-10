package pointersfunctions

import "fmt"

func update(num *int) {
	*num = 0
}
func display() *string {
	message := "Hi"
	return &message
}
func callByValue(num int) {
	num = 30
	fmt.Println(num) // 30
}
func callByReference(num *int) {
	*num = 10
	fmt.Println(*num) // 10
}
func Main() {
	n := 10
	update(&n)
	fmt.Println(n) //0
	rs := display()
	fmt.Println(*rs) //Hi
	var number int
	// passing value
	callByValue(number) //30
	// passing a reference (address)
	callByReference(&number) //10
	fmt.Println(number)      //10
}
