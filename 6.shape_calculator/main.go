// Problem: You have 3 shapes. Calculate total area of any shapes.

// Create interface:

// Go
// type Shape interface {
//     Area() float64
// }
// Implement for Rectangle, Circle, Triangle. Then write a function TotalArea(shapes []Shape) float64

// You will learn: Interfaces, implicit implementation - the core of Go OOP

// Bonus: Add Perimeter() to the interface.

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func TotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	return float64(totalArea)
}

func main() {

	rectangle := Rectangle{Width: 10, Length: 20}
	fmt.Println("Area of rectangle", rectangle.Area())
	fmt.Println("Perimeter of rectangle", rectangle.Perimeter())

	circle := Circle{Radius: 10}
	fmt.Println("Area of circle", circle.Area())
	fmt.Println("Perimeter of circle", circle.Perimeter())

	triangle := Triangle{A: 3, B: 4, C: 5}
	fmt.Println("Area of triangle", triangle.Area())
	fmt.Println("Perimeter of triangle", triangle.Perimeter())

	fmt.Println("TotalArea of all shapes", TotalArea([]Shape{rectangle, circle, triangle}))
}
