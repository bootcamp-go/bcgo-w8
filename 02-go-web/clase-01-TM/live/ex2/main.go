package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var err error

	// Encoder
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

	standardEncoder := json.NewEncoder(os.Stdout)

	if err = standardEncoder.Encode(data); err != nil {
		panic(err)
	}

	// Decoder
	var decodedData map[string]any
	toDecode := `{"age":30,"isMarried":true,"name":"John","shopList":["milk","apple","coffee"]}`

	standardDecoder := json.NewDecoder(strings.NewReader(toDecode))
	if err = standardDecoder.Decode(&decodedData); err != nil {
		panic(err)
	}

	fmt.Printf("%#v", decodedData)
}
