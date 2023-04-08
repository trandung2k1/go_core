package strings

import (
	"fmt"
	"strings"
)

func Strings() {
	var message1 = "Hello,"
	message2 := "Welcome to Programiz"

	fmt.Println(message1)
	fmt.Println(message2)

	//Using backtick
	message := `I love Go Programming`
	fmt.Println(message)

	//Access characters
	fmt.Printf("%c", message[0])
	stringLength := len(message)
	fmt.Println(stringLength) //21
	//Methods: Compare, Contains, Replaces, ToLower, ToUpper, Split, Join,

	string1 := "Program"
	string2 := "Program Pro"
	string3 := "Program"

	// compare strings
	fmt.Println(strings.Compare(string1, string2)) // -1
	fmt.Println(strings.Compare(string2, string3)) // 1
	fmt.Println(strings.Compare(string1, string3)) // 0

	//Create string from slice
	str := []string{"I", "love", "you"}
	fmt.Println(strings.Join(str, " ")) //I love you
	//Note: Go Strings are Immutable

}
