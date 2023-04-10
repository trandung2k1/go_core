package anonymousfunction

import "fmt"

func displayNumber() func() int {
	number := 10
	return func() int {
		number++
		return number
	}
}
func Example() {
	a := displayNumber()
	// Closure Go = Closure JavaScript
	fmt.Println("Called fun ", a()) //11
	fmt.Println("Called fun ", a()) //12
	fmt.Println("Called fun ", a()) //13

}
