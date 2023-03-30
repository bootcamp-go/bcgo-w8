package main

func increment(value *int) {
	*value++
}

func main() {
	var counter int

	println("counter:", counter)

	increment(&counter)

	println("counter:", counter)
	println("memory address of counter:", &counter)
}
