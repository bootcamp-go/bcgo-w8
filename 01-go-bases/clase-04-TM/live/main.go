package main

import (
	"errors"
	"fmt"
)

func main() {
	n1 := 10.0
	n2 := 0.0

	result, err := divide(n1, n2)
	if err != nil {
		if errors.Is(err, ErrZeroDivide) {
			fmt.Println("error zero division: ", err)
		} else if errors.Is(err, ErrInternal) {
			fmt.Println("error internal: ", err)
		} else {
			fmt.Println("error default: ", err)
		}
		return
	}

	fmt.Println("result:", result)
}

// package big calculator



// package calculator
var (
	ErrZeroDivide = errors.New("zero divide")
	ErrInternal   = errors.New("internal error")
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Value of B: %v is zero, %w", b, ErrZeroDivide)
	}

	return a / b, nil
}