package errors

import (
	"errors"
	"fmt"
)

func checkName(name string) error {
	//create new error
	newError := errors.New("Invalid Name")
	if name != "Dung" {
		return newError
	}
	return nil
}

func divide(num1, num2 int) error {
	if num2 == 0 {
		return fmt.Errorf("%d / %d\nCannot Divide a Number by zero", num1, num2)
	}
	return nil
}
func sum(a int, b int) int {
	fmt.Println(a)
	return a + b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("invalid division")
	}
	return a / b, nil
}
func Main() {
	//issues
	// for i := 0; i < 5; i++ {
	// 	rs := 20 / i
	// 	fmt.Println(rs) //panic: runtime error: integer divide by zero
	// }

	//fix issues: New() or Errof()
	// msg := "HI"
	// myError := errors.New("Wrong message")
	// if msg != "Hi" {
	// 	fmt.Println(myError)
	// 	fmt.Printf("%T", myError) //*errors.errorString
	// }

	name := "Dung"
	err := checkName(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid name") //Valid name
	}
	age := -14
	error := fmt.Errorf("%d is negative\n Age cant't be negative", age)
	if age < 0 {
		fmt.Println(error)
	} else {
		fmt.Println("Age: %d", age)
	}

	e := divide(4, 0)
	if err != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Valid division")
	}
	v, er := div(2, 2)
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println((v))
	}
}
