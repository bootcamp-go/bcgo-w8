package main

import "net/http"

func main() {
	var err error

	router := http.NewServeMux()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	if err = http.ListenAndServe(":1897", router); err != nil {
		panic(err)
	}
}
