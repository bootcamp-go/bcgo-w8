package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func main() {
	// server
	http.HandleFunc("/products", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}


// ___________________________________________________________________________
// package handlers (controller: comunicacion directa con el cliente)
func handler(w http.ResponseWriter, r *http.Request) {
	// auth: decodificamos el token y obtenemos el usuario
	user, err := ValidToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	// context
	ctx := context.WithValue(r.Context(), "user", user)


	// request (params, query, body, etc)
	// ...
	
	// process: items
	it, err := GetItems(ctx)
	if err != nil {
		if errors.Is(err, ErrUnauthorized) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// response
	b, err := json.Marshal(map[string]any{"message": "ok", "data": it})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK); w.Write(b)
}

// ___________________________________________________________________________
// package users
type User struct {
	ID   		int
	Name 		string
	AuthLevel 	int
}
func ValidToken(r *http.Request) (user *User, err error) {
	// verificamos que el token es valido (en este caso suponemos que lo es)
	
	// obtenemos el usuario
	user = &User{
		ID: 1,
		Name: "John",
		AuthLevel: 1,
	}
	return
}

// ___________________________________________________________________________
// package items
type Item struct {
	ID   		int
	Name 		string
	Price 		float64
}
var items = []Item{
	{ID: 1, Name: "Item 1", Price: 10.5},
	{ID: 2, Name: "Item 2", Price: 20.5},
	{ID: 3, Name: "Item 3", Price: 30.5},
}

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

func GetItems(ctx context.Context) (it []Item, err error) {
	// obtenemos el usuario
	user := ctx.Value("user").(*User)

	// verificamos el nivel de autorizacion
	if user.AuthLevel < 1 {
		err = ErrUnauthorized
		return
	}

	// filtramos los items
	it = append(it, items...)
	return
}
