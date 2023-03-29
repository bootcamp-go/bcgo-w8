package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// init client
	log.Println("[client] init client")
	client, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// listen server
	go func () {
		defer func () {
			if err := recover(); err != nil {
				log.Println("[log] error panic: ", err)
			}
			client.Close()
		}()

		for {
			// buffer: to save data
			buffer := make([]byte, 1024)

			// read conn message
			n, err := client.Read(buffer)
			if err != nil {
				if err.Error() == "EOF" {
					log.Println("[client] connection closed by server")
				} else {
					log.Println("[client] error read: ", err)
				}
				panic(err)
			}

			log.Println("[client] server message: ", string(buffer[:n]))
		}
	}()

	// io connection
	for {
		// std input with bufio
		log.Print("Enter message: ")

		var message string
		_, err := fmt.Scanln(&message)
		if err != nil {
			log.Println("[log] error read std-input: ", err)
			continue
		}

		// write conn message
		_, err = client.Write([]byte(message))
		if err != nil {
			log.Println("[client] error write: ", err)
			continue
		}
	}
}