package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// creamos un contexto con un timeout de 1 segundo
	ctxEmpty := context.Background()
	ctx, cancel := context.WithTimeout(ctxEmpty, 1*time.Second)
	defer cancel()

	// enviamos la peticion a la api
	// - caso 1: si tarda mas de 1 segundo, se llama automaticamente a cancel() y envia una señal al canal ctx.Done(),
	// 			 el cual reaccionaria primero que el canal ch
	// 			 en este caso, el canal ch no se ejecuta y se devuelve el error de timeout
	// 
	// - caso 2: la operacion de llamada a la api tarda menos de 1 segundo, en este caso el canal ch se ejecuta primero
	//			 y se devuelve el resultado
	result, err := sendRequest(ctx, "http://example.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

// Request a algun servicio externo con url
func sendRequest(ctx context.Context, url string) (result string, err error) {
	// async process
	ch := make(chan string)
	defer close(ch)

	go api(url, ch)

	// esperamos a que se complete la operacion o se cancele el contexto
	// caso 1: el resultado termina antes de que se cancele el contexto, devolviendo el resultado
	// caso 2: el contexto se cancela antes de que se complete la operacion, devolviendo un error de timeout o de que alguno lo cancelo desde otro lado
	// 		   ctx.Done recive una señal async cuando se cancela el contexto (manual) o cuando se acaba el timeout (automatico)
	for {
		select {
		case result = <-ch:
			return
		case <-ctx.Done():
			err = fmt.Errorf("timeout or cancelled: %w", ctx.Err())
			return
		}
	}
}

// Simula una operacion async de una api
func api(url string, ch chan string) {
	fmt.Println("request sent")
	time.Sleep(2 * time.Second)
	ch <- "response"
}