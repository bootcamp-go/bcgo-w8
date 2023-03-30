package main

import (
	"errors"
	"fmt"
)

func main() {
	var cc = calculatorLocal{}

	var calculator Calculator = &cc

	result, err := calculator.Add(1, 0)
	if err != nil {
		if errors.Is(err, ErrDivisionZero) {
			fmt.Println("number b is not possible to be zero")
			return
		}
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

// PACKAGE MATH
// Interface: calculadora
type Calculator interface {
	Add(a, b float64) (float64, error)
}
// -> errors
var (
	ErrDivisionZero = errors.New("division by zero")
)


// Implementación de la interfaz
type calculatorLocal struct{}
func (c *calculatorLocal) Add(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("number b is not possible to be zero: %w", ErrDivisionZero)
	}

	return a + b, nil
}