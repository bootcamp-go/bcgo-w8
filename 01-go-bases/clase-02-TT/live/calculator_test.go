package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Arrange.
	args := []float64{2, 2}
	expectedResult := 4.0

	// Act.
	obtainedResult := Add(args...)

	// Assert.
	assert.Equal(t, expectedResult, obtainedResult)
}

func TestDivide(t *testing.T) {
	t.Run("Divide two positive numbers", func(t *testing.T) {
		// Arrange.
		num1 := 10.0
		num2 := 2.0
		expectedResult := 5.0
		expectedOK := true

		// Act.
		obtainedResult, obtainedOK := Divide(num1, num2)

		// Assert.
		assert.Equal(t, expectedResult, obtainedResult)
		assert.Equal(t, expectedOK, obtainedOK)
	})

	t.Run("Cannot divide by zero", func(t *testing.T) {
		// Arrange.
		num1 := -10.0
		num2 := 0.0
		expectedResult := 0.0
		expectedOK := false

		// Act.
		obtainedResult, obtainedOK := Divide(num1, num2)

		// Assert.
		assert.Equal(t, expectedResult, obtainedResult)
		assert.Equal(t, expectedOK, obtainedOK)
	})
}
