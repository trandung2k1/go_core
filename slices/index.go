package slices

import "fmt"

func Slices() {
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	var numbers2 = []int{1, 2, 3}
	fmt.Println(numbers2)

	//Create from array
	var arrays = [6]int{1, 2, 3, 4, 5, 6}

	slicesNum := arrays[1:3]
	fmt.Println(slicesNum) //[2 3]

	//Methods: append, copy, Equal, len

	fmt.Print(slicesNum[0]) //2

	slicesNum[0] = 100
	fmt.Print(slicesNum[0]) //100

	//use make function
	data := make([]int, 5, 7) // 5 = length,
	// 7 is the capacity of the slice (maximum size up to which a slice can be extended)
	data[0] = 1
	data[1] = 2
	data[2] = 3
	data[3] = 4
	fmt.Println(data) //[1 2 3 4 0]

}
