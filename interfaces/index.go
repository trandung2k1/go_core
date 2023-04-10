package interfaces

import "fmt"

type Shape interface {
	area() float32
	perimeter() float32
}

type Rectangle struct {
	length, breadth float32
}

// use struct to implement area() of interface
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

func (r Rectangle) perimeter() float32 {
	return (r.length + r.breadth) * 2
}
func caculate(s Shape) {
	fmt.Println(s)             //{7, 4}
	fmt.Println(s.area())      // 28
	fmt.Println(s.perimeter()) //22
}

func Main() {
	rect := Rectangle{7, 4}
	caculate(rect) //28 , 22
	Example()
	EmptyInterface()
}
