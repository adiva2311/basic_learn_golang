package main

import (
	"fmt"
	"strconv"
)

type HasName interface {
	GetName() string
}

type Books struct {
	Name  string
	Stock int
}

func (b Books) GetName() string {
	result := b.Name + strconv.Itoa(b.Stock)
	return result
}

func ListBook(name HasName) {
	fmt.Println(name.GetName())
}

func main(){
	books := Books{
		Name: "Atomic Habits",
		Stock: 24,
	}
	fmt.Println(books)
}