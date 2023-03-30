package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}

func (circle Circle) Perimeter() float64 {
	return circle.Radius * 2 * math.Pi
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func GetInformation(shape Shape) string {
	return fmt.Sprintf(
		"Geometry shape has area %f and perimeter %f",
		shape.Area(),
		shape.Perimeter(),
	)
}

func main() {
	var myShape Shape = &Rectangle{
		Width:  10,
		Height: 5,
	}

	fmt.Println(GetInformation(myShape))

	shape := &Circle{
		Radius: 5,
	}

	fmt.Println(GetInformation(shape))
}
