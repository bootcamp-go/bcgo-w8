package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	IsMarried bool     `json:"isMarried"`
	ShopList  []string `json:"shopList,omitempty"`
	Password  string   `json:"-"`
}

func main() {
	// Crear un string JSON mediante un map.
	data := map[string]any{
		"name":      "John",
		"age":       30,
		"isMarried": true,
		"shopList": []string{
			"milk",
			"apple",
			"coffee",
		},
	}

	dataAsJSON, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data as map: %+v\n", data)
	fmt.Printf("Data as JSON: %s\n", dataAsJSON)

	// Crear un string JSON mediante una estructura.
	user := User{
		Name:      "John",
		Age:       30,
		IsMarried: true,
		ShopList:  []string{},
	}

	userAsJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User as struct: %+v\n", user)
	fmt.Printf("User as JSON: %s\n", userAsJSON)

	// Crear una estructura mediante un string JSON.
	var userFromJSON User

	if err = json.Unmarshal(userAsJSON, &userFromJSON); err != nil {
		panic(err)
	}

	fmt.Printf("User from JSON: %+v\n", userFromJSON)

	var userMapFromJSON map[string]any

	if err = json.Unmarshal(userAsJSON, &userMapFromJSON); err != nil {
		panic(err)
	}

	fmt.Printf("User map from JSON: %+v\n", userMapFromJSON)
}
