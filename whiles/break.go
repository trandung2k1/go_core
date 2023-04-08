package whiles

import "fmt"

func Break() {
	i := 0
	for {
		if i > 5 {
			break
		}
		fmt.Println(i) // 0 -> 5
		i++
	}
}
