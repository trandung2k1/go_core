package pointers

import "fmt"

func Example() {
	var num int
	var ptr *int

	num = 22
	fmt.Println("Address of num:", &num) //0xc0000a6098
	fmt.Println("Value of num:", num)    //22

	ptr = &num
	fmt.Println("\nAddress of pointer ptr:", ptr) //0xc0000a6098
	fmt.Println("Content of pointer ptr:", *ptr)  //22

	num = 11
	fmt.Println("\nAddress of pointer ptr:", ptr) //0xc0000a6098
	fmt.Println("Content of pointer ptr:", *ptr)  //11

	*ptr = 2
	fmt.Println("\nAddress of num:", &num) //0xc0000a6098
	fmt.Println("Value of num:", num)      // 2
}
