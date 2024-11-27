package main

import (
	"fmt"
)

// Merubah tipe data ke tipe data yang kita inginkan, biasa kita gunakan dengan interface kosong, interface{} or any

func random() interface{} {
	return 345
}

func main() {
	// Cara Pertama #1
	var value = random()
	// var valueString string = value.(string)
	// fmt.Println(valueString)

	var valueInt int = value.(int)
	fmt.Println(valueInt)

	// Cara Kedua Lebih dianjurkan
	var result = random()
	switch assertion := result.(type) {
	case string:
		fmt.Println("String", assertion)
	case int:
		fmt.Println("Integer", assertion)
	default:
		fmt.Println("Unknown", assertion)
	}
	
}