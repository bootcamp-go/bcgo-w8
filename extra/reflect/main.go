package main

import (
	"fmt"
	"reflect"
)

func main() {
	it := Item{Name: "item1", Price: 10.5, Quantity: 2}
	reflectCheck(&it)
}

type Item struct {
	Name     string  `json:"name" custom:"custom_name"`
	Price    float64 `json:"price" custom:"custom_price"`
	Quantity int     `json:"quantity" custom:"custom_quantity"`
}

// reflect check
func reflectCheck(ptr any) {
	dataV := reflect.ValueOf(ptr).Elem()
	dataT := reflect.TypeOf(ptr).Elem()

	// iterate over struct fields
	for i := 0; i < dataT.NumField(); i++ {
		fieldV := dataV.Field(i)
		fieldT := dataT.Field(i)

		// elements
		fieldName := fieldT.Name
		fieldType := fieldT.Type
		fieldTag := fieldT.Tag

		// print leaving black spaces between elements
		fmt.Println("struct")
		fmt.Printf("-name: %-10s -type: %-10s -tag: %-40s | -value: %v\n", fieldName, fieldType, fieldTag, fieldV)
		fmt.Printf("-tag_json: %-10s -tag_custom: %-10s\n", fieldTag.Get("json"), fieldTag.Get("custom"))

		fmt.Println()
	}
}