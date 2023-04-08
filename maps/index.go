package maps

import "fmt"

func Map() {
	// subjectMarks := map[string]float32{"Go": 85, "Java": 80}
	var subjectMarks = map[string]float32{"Go": 85, "Java": 80}
	fmt.Println(subjectMarks) //map[Go:85 Java:80]
	//Access
	fmt.Println(subjectMarks["Go"]) //85

	//Change
	subjectMarks["Java"] = 100
	fmt.Println(subjectMarks["Java"]) //100

	//Add
	subjectMarks["C"] = 140
	fmt.Println(subjectMarks) //map[C:140 Go:85 Java:100]

	//Delete

	personAge := map[string]int{"Hermione": 21, "Harry": 20, "John": 25}
	delete(personAge, "John")

	fmt.Println(len(personAge)) //2

	//Loop
	for i, v := range personAge {
		fmt.Println(i, v)
	}

	for _, v := range personAge {
		fmt.Println(v)
	}

	//Create map use make function

	data := make(map[int]string)
	data[1] = "Harry"
	data[2] = "Alice"

	fmt.Println(data) //map[1:Harry 2:Alice]

}
