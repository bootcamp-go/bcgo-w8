package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a http server
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}
	defer server.Close()

	// Start the server
	log.Println("[server] Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("[server] Unable to start server - ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// recover from panic
	defer func() {
		if r := recover(); r != nil {
			log.Println("[server] Recovered from a panic - ", r)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	// io communication (api)
	switch r.Method {
	case http.MethodGet:
		// send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello World"}`))
	default:
		panic("Method not allowed")
	}
}