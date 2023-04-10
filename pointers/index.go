package pointers

import "fmt"

func Pointers() {
	var x int = 1
	fmt.Println("Address x", &x) //0xc00001c098
	var ptr *int = &x
	fmt.Println(ptr)  //0xc00001c098
	fmt.Println(*ptr) //1
	Example()
	var str *string
	fmt.Println(str) // <nil>
	var s = new(int)
	*s = 10
	fmt.Println(*s) //10
}
