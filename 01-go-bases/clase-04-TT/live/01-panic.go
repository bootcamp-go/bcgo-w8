package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Task: execute a function to open a file
	data := ReadData("not-exists.txt")
	println(data)

	// native panic in go
	// Task: add a key to a map
	var hashmap map[string]interface{}
	hashmap["key"] = "value"
	hashmap["name"] = "valeria"

	// Task: read an slice from a range of indexes
	var sl = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sl[5:11])

	// Task: inicialize a service and execute a function
	// api := APIImpl{}
	sv := Service{api: nil}

	weather := sv.api.GetWeather()
	println(weather)
}

type Service struct {
	// ...
	api API
}
type API interface {
	GetWeather() string
}
type APIImpl struct {}
func (a *APIImpl) GetWeather() string {
	return "sunny"
}

func ReadData(filename string) string {
	f1, err := os.Open(filename)
	if err != nil {
		panic("Error opening file: " + err.Error())
	}

	// type req struct {
	// 	nombre string
	// }
	// var r req

	// d := json.NewDecoder(f1)
	// d.Decode(&r)

	text, err := io.ReadAll(f1)
	if err != nil {
		panic(err)
	}

	return string(text)
}