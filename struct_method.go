package main

import "fmt"

type Products struct {
	Name  string
	Code  string
	Stock int
}

func (product Products) listProduct(name string) {
	fmt.Println(name,", membeli sebuah laptop merk ",product.Name,"dengan kode ",product.Code,", maka sisa stocknya adalah ",product.Stock)
}

func main() {
	laptop := Products{
		Name: "Asus",
		Code: "X431SA",
		Stock: 23,
	}
	laptop.listProduct("Joko")
}