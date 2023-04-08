package structs

import "fmt"

type Person struct {
	name string
	age  int
}

type Rectangle func(int, int) int

type rectanglePara struct {
	length  int
	breadth int
	color   string
	rect    Rectangle
}

func Struct() {
	var person1 Person
	person1.name = "John"
	person1.age = 21
	fmt.Println(person1)

	person := Person{
		name: "Dung",
		age:  22,
	}
	fmt.Println(person)

	fmt.Println(person.age) //22

	fmt.Printf("%T", person) //structs.Person

	result := rectanglePara{
		length:  10,
		breadth: 20,
		color:   "Red",
		rect: func(length, breadth int) int {
			return length * breadth
		},
	}
	fmt.Println()
	fmt.Println(result.color)                               //Red
	fmt.Println(result.rect(result.length, result.breadth)) //200
}
