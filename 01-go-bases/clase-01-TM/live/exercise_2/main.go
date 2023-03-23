package main

import "fmt"

func main() {
	// Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura,
	// humedad y presión de donde te encuentres.
	var (
		temperature float64 = 26.2
		humidity    float64 = 78.7
		pressure    float64 = 1013.0
	)

	// Imprimí los valores de las variables en consola.
	fmt.Println("Temperature:", temperature)
	fmt.Println("Humidity:", humidity)
	fmt.Println("Pressure:", pressure)
}
