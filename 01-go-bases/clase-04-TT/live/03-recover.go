package main

import "fmt"

func main() {
	// create a function to check if number is odd, if not, panic
	n := 10

	IsOdd(n)

	println("done")
}

func IsOdd(n int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("program has recover", err)
		}
	}()

	if n%2 == 0 {
		panic("number is even")
	}

	println("number is odd")
}