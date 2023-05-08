package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple Test
func TestCalcular_Add_SameNumbers(t *testing.T) {
	// arrange
	n := 5
	c := NewCalculator(1)

	exp := 10

	// act
	res, err := c.Add(n, n)

	// assert
	assert.Equal(t, exp, res)
	assert.NoError(t, err)
	assert.ErrorIs(t, err, nil)
}
func TestCalcular_Add_UnavailableOperation(t *testing.T) {
	// arrange
	n := 5
	c := NewCalculator(0)

	// act
	res, err := c.Add(n, n)

	// assert
	assert.Equal(t, 0, res)
	assert.ErrorIs(t, err, ErrInternal)
	assert.EqualError(t, err, "internal error: unavailable operation")
}

// Subtests
func TestCalcular_Add(t *testing.T) {
	// succeed cases
	t.Run("same number", func(t *testing.T) {
		// arrange
		n := 5
		c := NewCalculator(2)

		// act
		res, err := c.Add(n, n)

		// assert
		assert.Equal(t, 10, res)
		assert.NoError(t, err)
		assert.Equal(t, 1, c.counter)
	})

	t.Run("different numbers", func(t *testing.T) {
		// arrange
		n1 := 5
		n2 := 2

		exp := 7

		c := NewCalculator(2)

		// act
		res, err := c.Add(n1, n2)

		// assert
		assert.Equal(t, exp, res)
		assert.NoError(t, err)
		assert.Equal(t, 1, c.counter)
	})

	t.Run("negative numbers", func(t *testing.T) {
		// arrange
		n1 := -5
		n2 := -2

		exp := -7

		c := NewCalculator(2)

		// act
		res, err := c.Add(n1, n2)

		// assert
		assert.Equal(t, exp, res)
		assert.NoError(t, err)
		assert.Equal(t, 1, c.counter)
	})
}

// Table Driven Tests
func TestCalculator_Sum_TDT(t *testing.T) {
	type input struct {n1, n2 int}
	type output struct {res int; err error}
	
	type testCase struct {
		// base
		name   string
		input  input
		output output
		// process
		setCalculator func(c *Calculator)
	}
	cases := []testCase{
		// succeed cases
		{
			name: "same number",
			input: input{n1: 5, n2: 5},
			output: output{res: 10, err: nil},
			setCalculator: func(c *Calculator) {c.limiter = 2},
		},
		{
			name: "different numbers",
			input: input{n1: 5, n2: 2},
			output: output{res: 7, err: nil},
			setCalculator: func(c *Calculator) {c.limiter = 2},
		},
		{
			name: "negative numbers",
			input: input{n1: -5, n2: -2},
			output: output{res: -7, err: nil},
			setCalculator: func(c *Calculator) {c.limiter = 2},
		},
		// failure cases
		{
			name: "unavailable operation",
			input: input{n1: 5, n2: 5},
			output: output{res: 0, err: ErrInternal},
			setCalculator: func(c *Calculator) {c.limiter = 0},
		},
	}

	// run test cases
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			c := NewCalculator(0)
			tc.setCalculator(c)

			// act
			res, err := c.Add(tc.input.n1, tc.input.n2)

			// assert
			assert.ErrorIs(t, err, tc.output.err)
			assert.Equal(t, tc.output.res, res)
		})
	}
}