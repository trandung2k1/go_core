package closure

import "fmt"

//nested function

func greet(name string) {
	var displayName = func() {
		fmt.Println("Hi", name)
	}
	displayName()
}



func Example() {
	greet("Dung")
	
}
