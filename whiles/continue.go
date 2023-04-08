package whiles

import "fmt"

func Continue() {
	i := 1

	for i < 10 {
		if i%2 == 0 {
			i++
			continue
		}
		fmt.Println(i) // 1 3 5 7 9
		i++
	}
}
