package datatypes

import "fmt"

func DataTypes() {
	//int: int/uint(32bit or 64bit), 8bit = 1byte, int8/uint8, .. 16, 32, 64 bit

	var id int
	var age uint
	id = 1
	age = 21
	fmt.Printf("\n%d", id)  //1
	fmt.Printf("\n%d", age) //21

	//float: float32 and float64
	salary := 5676.3           // default 64 bit
	fmt.Printf("\n%f", salary) //5676.300000

	var x float32 = 50000.503882901
	fmt.Printf("\n%f", x) //50000.503906
	fmt.Println()
	fmt.Println(x) //50000.504

	//string
	var message string = "Hello, world"
	msg := `Hi`
	fmt.Print(message) //Hello, world
	fmt.Println()
	fmt.Print(msg) //Hi

	var c rune = 'A'
	fmt.Println()
	fmt.Println(c) // 65

	//boolean
	var isLoading bool = false
	fmt.Println(isLoading) //false

	var b byte = 1
	fmt.Println(b)
}
