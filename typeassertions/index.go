package typeassertions

import "fmt"

func Main() {
	var a interface{}
	a = 10
	interfaceValue := a.(int)
	fmt.Println(interfaceValue) //10

	value, isValid := a.(int)

	fmt.Println(value, isValid) //10 true

	v, i := a.(string)
	fmt.Println(v, i) // i = false
	var b interface{}
	b = "Golang"
	vl := b.(int)
	fmt.Println(vl) //panic: interface conversion: interface {} is string, not int
}
