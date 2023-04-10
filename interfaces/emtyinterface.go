package interfaces

import "fmt"

func displayValue(i interface{}) {
	fmt.Println(i)
}

// Arguments
func displayData(args ...interface{}) {
	fmt.Println(args)    //[1 true Hi]
	fmt.Println(args...) //1 true Hi
}
func EmptyInterface() {
	var a interface{}
	fmt.Println(a) // <nil>
	a = "Hi"
	a = 1
	fmt.Println(a) //1
	b := true
	c := "Hi"
	displayValue(a) //1
	displayValue(b) //true
	displayValue(c) //Hi
	displayData(1, true, "Hi")
}
