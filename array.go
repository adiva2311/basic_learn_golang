package main

import "fmt"

func main() {
	var fruits = [...]string{
		"apple",
		"orange",
		"grape",
		"lemon",
		"lotus",
	}

	for _, fruit := range fruits {
		fmt.Printf("%s is fruit \n", fruit)
	}

}
