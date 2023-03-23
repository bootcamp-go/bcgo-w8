package main

import "fmt"

func main() {
	IfStatement()
	SwitchStatement()
}

func IfStatement() {
	// Task: crear un programa que chequee la edad de una persona
	// si la persona es mayor de 18 años, imprimir "puede votar"
	// si la persona es mayor de 16 años, imprimir "puede conducir"
	// si la persona es mayor de 14 años, imprimir "puede ver películas clasificadas PG-13"
	// si la persona es menor de 14 años, imprimir "no puede hacer nada"

	// Solution:
	var edad int
	edad = 20

	// caso 1 (exclusivo)
	// if edad >= 18 {
	// 	fmt.Println("puede votar")
	// } else if edad >= 16 {
	// 	fmt.Println("puede conducir")
	// } else if edad >= 14 {
	// 	fmt.Println("puede ver películas clasificadas PG-13")
	// } else {
	// 	fmt.Println("no puede hacer nada")
	// }


	// caso 2
	// validacion
	if edad <= 14 {
		fmt.Println("no puede hacer nada")
		return
	}

	// proceso (inclusivo)
	if edad >= 18 {
		fmt.Println("puede votar")
	}
	if edad >= 16 {
		fmt.Println("puede conducir")
	}
	if edad >= 14 {
		fmt.Println("puede ver películas clasificadas PG-13")
	}
}

func SwitchStatement() {
	// Task: crear un programa que chequee el puesto de trabajo de una persona
	// - caso de ser ceo: imprimir "puede hacer cualquier cosa"
	// - caso de ser manager: imprimir "puede manejar"
	// - caso de ser supervisor: imprimir "puede supervisar"
	// - caso de ser employee: imprimir "puede hacer su trabajo"

	// Solucion:
	edad := 20

	// validacion de numeros negativos
	if edad < 0 {
		fmt.Println("edad no puede ser negativa")
		return
	}

	// proceso
	switch {
	case edad >= 18:
		fmt.Println("puede votar")
		fallthrough													// fallthrough es para que se ejecute el siguiente caso sin importar si se cumple o no
	case edad >= 16:
		fmt.Println("puede conducir")
		fallthrough
	case edad >= 14:
		fmt.Println("puede ver películas clasificadas PG-13")
	default:
		fmt.Println("no puede hacer nada")
	}
}