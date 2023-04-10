package interfaces

import "fmt"

type Shape1 interface {
	area() float32
}
type Rectangle1 struct {
	length, breadth float32
}

func (r Rectangle1) area() float32 {
	return r.length * r.breadth
}

type Triangle struct {
	base, height float32
}

func (t Triangle) area() float32 {
	return 0.5 * t.base * t.height
}
func calculate1(s Shape1) float32 {
	return s.area()
}

//Note:  implement all methods of interface
func Example() {
	r := Rectangle1{7, 4}
	t := Triangle{8, 12}

	// call calculate() with struct variable rect
	rectangleArea := calculate1(r)
	fmt.Println("Area of Rectangle:", rectangleArea) //Area of Rectangle: 28

	triangleArea := calculate1(t)
	fmt.Println("Area of Triangle:", triangleArea) //Area of Triangle: 48
}
