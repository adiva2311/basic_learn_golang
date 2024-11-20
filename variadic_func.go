package main

import "fmt"

func sumAll(numbs ...int) {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	fmt.Println(result)
}

func main() {
	sumAll(10, 20, 30)
}