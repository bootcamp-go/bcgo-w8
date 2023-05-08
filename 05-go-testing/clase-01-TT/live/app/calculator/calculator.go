package calculator

import (
	"errors"
	"fmt"
)

var (
	ErrInternal     = errors.New("internal error")
	ErrZeroDivision = errors.New("division by zero")
)

func NewCalculator(limiter int) *Calculator {
	return &Calculator{limiter: limiter}
}

// Calculator is a simple calculator that can add and divide numbers.
type Calculator struct {
	limiter, counter int
}
func (c *Calculator) Add(a, b int) (result int, err error) {
	if err = c.available(); err != nil {
		return
	}

	result = a + b
	c.counter++
	return
}
func (c *Calculator) Divide(a, b int) (result int, err error) {
	if err = c.available(); err != nil {
		return
	}
	
	if b == 0 {
		err = fmt.Errorf("%w: %d / %d", ErrZeroDivision, a, b)
		return
	}
	result = a / b
	c.counter++
	return
}

func (c *Calculator) Exponential(a int) (result int, err error) {
	if err = c.available(); err != nil {
		return
	}
	result = a * a
	c.counter++
	return
}

// extra feature
func (c *Calculator) available() (err error) {
	if c.counter >= c.limiter {
		err = fmt.Errorf("%w: unavailable operation", ErrInternal)
	}
	return
}