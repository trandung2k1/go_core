package goroutine

import (
	"fmt"
	"time"
)

// func display(message string) {
// 	fmt.Println(message)

// }

func display(message string) {
	for {
		fmt.Println(message)
		time.Sleep(time.Second * 1)
	}
}
func Main() {
	go display("HI") // vs main()
	display("Hello") // Hello
}
