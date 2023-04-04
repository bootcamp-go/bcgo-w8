package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "usuario", "jorge")
	ctx = context.WithValue(ctx, "carrito", Carrito{Productos: []string{"camisa", "pantalon"}})

	go LogUsers(ctx)

	ProcesarCarrito(ctx)
}

// tareas
func LogUsers(ctx context.Context) {
	usuario := ctx.Value("usuario").(string)
	fmt.Println(usuario)
}

type Carrito struct {
	Productos []string
}
func ProcesarCarrito(ctx context.Context) {
	user := ctx.Value("usuario").(string)
	carrito := ctx.Value("carrito").(Carrito)

	// procesar carrito
	fmt.Println("procesando carrito de", user)
	for _, producto := range carrito.Productos {
		fmt.Println("-", producto)
	}

	fmt.Println("carrito procesado")
}