package arrays

import "fmt"

func Arrays() {
	var arrays = [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arrays)) //5
	fmt.Println(arrays[0])   //1

	//
	var names = [...]string{"Lan"}
	fmt.Println(names)

	fruits := [2]string{"Mango", "Orange"}
	fmt.Println(fruits) //[Mango Orange]

	strings := [...]string{"Hello", "Hi"}
	fmt.Println(strings) //[Hello Hi]

	//Loop array

	for i := 0; i < len(arrays); i++ {
		fmt.Println(arrays[i]) //1 2 3 4 5
	}

	for i, v := range arrays {
		fmt.Println(i, v) //
	}

	var cars [5]string
	cars[0] = "Mercedes"
	fmt.Println(cars)
	
	for i, v := range cars {
		fmt.Println(i, v) //
	}

}
