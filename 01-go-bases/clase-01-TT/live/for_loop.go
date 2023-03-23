package main

import "fmt"

func main() {
	Counter()
	IterateArray()
	IterateSlice()
	IterateMap()
	Iterate()
}

func Counter() {
	// Crear un incrementador de a 2
	// - solo se debe repitir 5 veces
	// var lim int = 5
	// var num int = 10

	// Solution
	var lim = 5
	num := 10

	// __________________________________________________________
	// s1: bucle infinito (break)
	// cc := 0
	// for {
	// 	// condicion de salida
	// 	if cc == lim {
	// 		break
	// 	}
	// 	// incremento
	// 	num += 2
	// 	// incremento del contador
	// 	cc++
	// }
	
	// __________________________________________________________
	// s2: bucle infinito (condicion incluida)
	// for cc < lim {
	// 	num += 2
	// 	cc++
	// }

	// __________________________________________________________
	// s3: bucle con rango
	for i:=0; i<lim; i++ {
		num += 2
	}

	fmt.Println("num: ", num)
}

func IterateArray() {
	// Crear un array de items
	// - mostrar el primer item
	// - mostrar el último item

	// Solution
	frutas := [4]string{"banana", "manzana", "pera", "uva"}
	fmt.Println("[info] len: ", len(frutas), "cap: ", cap(frutas))

	// primer item
	fmt.Println("primer item: ", frutas[0])

	// ultimo item
	fmt.Println("ultimo item: ", frutas[len(frutas)-1])
}

func IterateSlice() {
	// Crear un slice de items
	// - mostrar el primer item
	// - mostrar el último item
	// - agregar un item al slice
	// - verificar si el item existe

	// Solution
	// frutas := make([]string, 4)										// len: 4, cap: 4 (capacidad = longitud inicial  -> frutas = ["" "" "" ""]
	// frutas := make([]string, 4, 8)									// len: 4, cap: 8 , en este caso la capacidad es especificada
	frutas := []string{"banana", "manzana", "pera", "uva"}			// len: 4, cap: 4 (capacidad = longitud inicial)
	fmt.Println("[info] len: ", len(frutas), "cap: ", cap(frutas))

	// mostrar primer
	fmt.Println("primer item: ", frutas[0])
	// mostrar ultimo
	fmt.Println("ultimo item", frutas[len(frutas)-1])
	// agregar un item al slice
	frutas = append(frutas, "naranja")

	// mostrar nueva capacidad y longitud
	fmt.Println("[info] len: ", len(frutas), "cap: ", cap(frutas))


	// iterar
	fmt.Println("todos los items")
	for ix, item := range frutas {
		fmt.Println(ix, item)
	}

	// nuevo slice con make
	fmt.Println("\nNuevo Slice")
	sl := make([]int, 0, 5)       // len: 0, cap: 5
	fmt.Println("sl: ", sl)
	fmt.Println("len: ", len(sl), "cap: ", cap(sl))

	sl = append(sl, 1, 2, 3, 4, 5, 6)
	fmt.Println("sl: ", sl)
	fmt.Println("len: ", len(sl), "cap: ", cap(sl))

	// new
	sl2 := new([]int)

	fmt.Println("sl2: ", sl2)
}

func IterateMap() {
	// Crear un registro de personas y sus items
	// - mostrar la persona y sus items
	// - agregar una persona y sus items
	// - mostrar el registro de todas las personas y sus items

	// Solution
	var registro = make(map[string][]string)
	registro["juan"] = []string{"banana", "manzana", "pera", "uva"}

	fmt.Println("registro: ", registro)

	// mostrar persona
	items, ok := registro["juan"]
	if ok {
		fmt.Println("- items de juan: ", items)
	}

	// borrar persona
	fmt.Println("se elimina a juan")
	delete(registro, "juan")
	fmt.Println("registro: ", registro)

	// iterate
	fmt.Println("se agrega mas personas")
	registro["pedro"] = []string{"banana", "manzana", "pera", "uva"}
	registro["maria"] = []string{"naraja", "limon", "sandia", "melon"}

	for persona, items := range registro {
		fmt.Println("persona: ", persona)
		fmt.Println("items: ", items)
	}
}

func Iterate() {
	// A partir de un slice de items, contar cuántas veces se repite cada item y mostrarlo en un registro
	var frutas = []string{"banana", "manzana", "pera", "uva", "banana", "manzana", "pera", "uva", "banana", "manzana"}
	
	var hashmap = make(map[string]int)

	for _, item := range frutas {
		hashmap[item]++
		// if _, ok := hashmap[item]; ok {
		// 	hashmap[item] += 1
		// } else {
		// 	hashmap[item] = 1
		// }
	}

	fmt.Println("hashmap: ", hashmap)
}