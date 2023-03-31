package main

import (
	"context"
	"fmt"
)

// El context se utiliza para propagar información a través de la llamada de funciones. Esto es especialmente util tambien durante la ejecución de gorutinas.
// Casos de uso:
// - guardar información de la request	 -> context.WithValue
// - cancelar una operación				 -> context.WithCancel
// - cancelar una operación con deadline -> context.WithDeadline  (fecha especifica)
// - cancelar una operación con timeout  -> context.WithTimeout   (lapso de tiempo)
func main() {
	// context vacio
	var ctx context.Context
	// -> 1
	ctx = context.Background()
	// -> 2
	ctx = context.TODO()


	// context with value
	ctx = context.WithValue(ctx, "key", "value")
	fmt.Println("context value:", ctx.Value("key"))
}