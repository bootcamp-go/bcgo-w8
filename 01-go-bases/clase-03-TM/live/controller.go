package main

import "fmt"

// constructor
func NewCalculator(lim int) *Calculator {
	return &Calculator{lim: lim}
}

// controller
type Calculator struct {
	// config
	lim        int
	operations int
}

func (c *Calculator) available() bool {
	return c.operations < c.lim
}

func (c *Calculator) Sum(a, b int) int {
	if !c.available() {
		return 0
	}

	c.operations++
	return a + b
}
func (c *Calculator) Sub(a, b int) int {
	if !c.available() {
		return 0
	}

	c.operations++
	return a - b
}
func (c *Calculator) Mul(a, b int) int {
	if !c.available() {
		return 0
	}

	c.operations++
	return a * b
}
func (c *Calculator) Div(a, b int) int {
	if !c.available() {
		return 0
	}

	c.operations++
	return a / b
}

func main() {
	// init
	calc := NewCalculator(3)

	// use
	fmt.Println("sum:", calc.Sum(1, 2))
	fmt.Println("sum:", calc.Sum(1, 2))
	fmt.Println("sum:", calc.Sum(1, 2))

	fmt.Println("sum:", calc.Sum(1, 2))
}


// func Sum(a, b int) int {
// 	return a + b
// }

// func Sub(a, b int) int {
// 	return a - b
// }

// func Mul(a, b int) int {
// 	return a * b
// }

// func Div(a, b int) int {
// 	return a / b
// }