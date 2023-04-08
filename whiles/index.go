package whiles

import "fmt"

func Whiles() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	//Do while
	y := 1
	for {
		if y > 5 {
			break
		}
		fmt.Println(y)
		y++
	}
}
