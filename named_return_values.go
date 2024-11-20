package main

import "fmt"

func getCompleteName() (first, mid, last string, age int) {
	first = "Adiva"
	mid = "Nursuandy"
	last = "Ritonga"
	age = 24

	return first, mid, last, age
}

func main() {
	a, b, c, d := getCompleteName()
	fmt.Println(a, b, c, d)
}