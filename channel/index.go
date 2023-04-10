package channel

import "fmt"

func channelData(number chan int, message chan string) {
	number <- 15
	message <- "Hello, world!"
}
func Main() {
	name := make(chan string)
	number := make(chan int)
	message := make(chan string)
	fmt.Printf("%T\n", name) //chan string
	fmt.Printf("%v\n", name) //0xc000086060

	//send data
	go channelData(number, message)
	fmt.Println("Channel Data:", <-number)
	fmt.Println("Channel Data:", <-message)

}
