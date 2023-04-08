package typecastings

import "fmt"

func TypeCastings() {

	//convert float -> int
	var floatValue float32 = 5.45
	var intValue int = int(floatValue)
	fmt.Printf("%d\n", intValue) // 5
	fmt.Printf("%T\n", intValue) // int

	//Character
	a := 'a'
	fmt.Printf("%c\n", a) // 'a'

	//Convert int -> float
	var x int = 1
	var y float64 = float64(x)
	fmt.Printf("%f\n", y) //1.000000
	fmt.Printf("%T\n", y) //float64

	// var num int = 1.5 // error: int is a signed integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, int32.

	//Example

	var number1 int = 20
	var number2 float32 = 5.7
	var sum float32 = float32(number1) + number2
	fmt.Println(sum) // 25.7
}
