package main

import (
	"log"
	"net"
)

func main() {
	// init server
	log.Println("[server] init server")
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer server.Close()

	// init listener
	for {
		// client connection
		conn, err := server.Accept()
		if err != nil {
			log.Println("[server] error connection: ", err)
			continue
		}

		// handles connection (read request, write response)
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func () {
		if err := recover(); err != nil {
			log.Println("[server] panic: ", err)
		}
		conn.Close()
	}()

	// io connection (stream)
	for {
		// buffer: to save data
		buffer := make([]byte, 1024)

		// read conn message [request]
		n, err := conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("[server] connection closed by client")
			} else {
				log.Println("[server] error read: ", err)
			} 
				
			break
		}

		log.Println("[server] client message: ", string(buffer[:n]))

		// write conn message [response - api]
		_, err = conn.Write([]byte("Hello client!"))
		if err != nil {
			log.Println("[server] error write: ", err)
			continue
		}
	}
}