package main

import "fmt"

// Cara kedua adalah menggunakan Type Declaration

type Filter func (string) string

func getHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing"{
		return "..."
	} else {
		return name
	}
}

func main(){
	// CARA PERTAMA (1)
	getHelloWithFilter("Adiva", spamFilter)

	// CARA KEDUA (2)
	// Jika function MEMPUNYAI (), maka function akan dipanggil
	// Jika Function TIDAK MEMPUNYAI (), maka function akan dianggap sebagai variabel atau value
	filter := spamFilter
	getHelloWithFilter("Anjing", filter)
}