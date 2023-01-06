package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}
type rectangle struct {
	length float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r rectangle) area() float64 {
	return r.length * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	fmt.Printf("Enter radius of Circle: ")
	var radius float64
	fmt.Scanf("%f", &radius)

	cir := circle{radius}
	fmt.Printf("Area of Circle: ", cir.area())
	fmt.Println()
	fmt.Printf("Perimeter of Circle: ", cir.perimeter())
	fmt.Println()

	fmt.Printf("Enter Length and Width: ")

	var length, width float64
	fmt.Scanf("%f %f", &length, &width)
	rec := rectangle{length, width}
	fmt.Printf("Area of Rectangle: ", rec.area())
	fmt.Println()
	fmt.Printf("Perimeter of Rectangle: ", rec.perimeter())
}
