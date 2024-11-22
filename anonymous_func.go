package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're Block", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	registerUser("Adiva", func(s string) bool {
		return s == "Anjing"
	})
}