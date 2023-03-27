package main

import (
	"errors"
	"fmt"
)

const (
	suma           = "+"
	resta          = "-"
	division       = "/"
	multiplicacion = "*"
)

func main() {
	//inspeccionarVariable(0)
	//scope()
	//elipsisSimple(1,2,3,4,5,6)
	//elipsisMultiple("Alan","soy capo",1,2,2,2)
	//eligeOp("+",1,2,3)
	//multipleRetorno(25,2)
	//retornoValoresNombrados(3,2)
	//funcionDevuelveFuncion("gato")(10)
}

func inspeccionarVariable(numero int) {
	if numero < 0 {
		fmt.Println("Negativo")
	} else if numero > 0 {
		fmt.Println("Positivo")
	} else {
		fmt.Println("cero")
	}
}

func scope() {
	for i := 0; i < 3; i++ {
		fmt.Printf("i: %v\n", i)
		for j := 0; j < 3; j++ {
			var num = 3
			fmt.Printf("num: %v\n", num)
			//fmt.Printf("%v",i)
		}
		//fmt.Printf(num)
	}
	//fmt.Printf(i)
}

func elipsisSimple(valores ...float64) float64 {
	var resultado float64
	for _, valor := range valores {
		resultado += valor
	}
	return resultado
}

func elipsisMultiple(valor1 string, valor2 string, valores ...float64) {
	fmt.Println(valor1)
	fmt.Println(valor2)
	valoresCopy := make([]float64, len(valores))
	copy(valoresCopy, valores)
	fmt.Printf("%v ", valoresCopy)
}

func opsuma(num1, num2 float64) float64 {
	return num1 + num2
}
func opresta(num1, num2 float64) float64 {
	return num1 - num2
}
func opmultiplicacion(num1, num2 float64) float64 {
	return num1 * num2
}
func opdivision(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}

func haceLasCuentas(valores []float64, operacion func(valor1, valor2 float64) float64) float64 {
	var resul float64
	for _, valor := range valores {
		resul = operacion(resul, valor)
	}
	return resul
}

func eligeOp(operador string, valores ...float64) float64 {
	switch operador {
	case suma:
		return haceLasCuentas(valores, opsuma)
	case resta:
		return haceLasCuentas(valores, opresta)
	case division:
		return haceLasCuentas(valores, opdivision)
	case multiplicacion:
		return haceLasCuentas(valores, opmultiplicacion)
	}
	return 0
}

func multipleRetorno(num1, num2 int) (string, error) {
	if num2 == 0 {
		return "", errors.New("No se puede dividir por cero")
	}
	if num1%num2 == 0 {
		return "es divisible", nil
	} else {
		return "no es divisible", nil
	}
}

func retornoValoresNombrados(num1, num2 float64) (esMayor float64, mensaje string) {
	if num1 > num2 {
		esMayor = num1
		mensaje = "El primer numero es mayor"
	} else if num2 > num1 {
		esMayor = num2
		mensaje = "el segundo numero es mayor"
	} else {
		esMayor = 0
		mensaje = "son iguales"
	}
	return
}

func funcionDevuelveFuncion(mascota string) func(peso float64) float64 {
	switch mascota {
	case "perro":
		return calculoPerro
	case "gato":
		return calculoGato
	}
	return nil
}

func calculoPerro(peso float64) float64 {
	return peso * 0.3
}
func calculoGato(peso float64) float64 {
	return peso * 0.2
}
