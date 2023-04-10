package pointersstructs

import "fmt"

type Person struct {
	name string
	age  int
}

func Main() {
	var person = Person{"Dung", 21}
	var ptr *Person
	ptr = &person
	fmt.Println(ptr)  //&{Dung 21}
	fmt.Println(*ptr) //{Dung 21}

	//Access

	fmt.Println(ptr.age)    //21
	fmt.Println((*ptr).age) //21

	//Change value
	// *&ptr.age = 22
	ptr.age = 22

	fmt.Println(person) //{Dung 22}
}
