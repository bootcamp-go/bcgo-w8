package calculator

func Add(numbers ...float64) (result float64) {
	for i := range numbers {
		result += numbers[i]
	}
	return
}

func Divide(num1, num2 float64) (result float64, ok bool) {
	if num2 == 0 {
		return
	}

	result = num1 / num2
	ok = true
	return
}
