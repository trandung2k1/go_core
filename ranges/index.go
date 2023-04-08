package ranges

import "fmt"

func Ranges() {

	//Use array
	numbers := [5]int{1, 2, 3, 4, 5}
	for index, item := range numbers {
		fmt.Println(index, " - ", item)
	}

	//Use string
	str := "Hello, world"
	for i, c := range str {
		fmt.Printf("%d, %c\n", i, c)
	}

	//Use map

	subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}
	for item := range subjectMarks {
		fmt.Println(item)
	}

	for i, value := range subjectMarks {
		fmt.Println(i, value)
	}

	fmt.Printf("%T", subjectMarks) //map[string]float32
}
