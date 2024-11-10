package main

import "fmt"

func main() {
	var fruitsA = []string{"apel", "anggur", "jeruk"}    //Slice
	var fruitsB = [3]string{"apel", "anggur", "jeruk"}   //Array
	var fruitsC = [...]string{"apel", "anggur", "jeruk"} //Array

	for _, fruit := range fruitsA {
		fmt.Printf("%s is fruit \n", fruit)
	}
	for _, fruit := range fruitsB {
		fmt.Printf("%s is fruit \n", fruit)
	}
	for _, fruit := range fruitsC {
		fmt.Printf("%s is fruit \n", fruit)
	}

}
