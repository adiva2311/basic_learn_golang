package main

import "fmt"

/*
type NAMA_STRUCT struct{
	Nama_field Type_data
}
*/
type Customer struct{
	Name, Address string
	Age int
}

func main(){
	// Cara Deklari Struct 1
	var adiva Customer
	adiva.Name = "Adiva Nursuandy"
	adiva.Address = "Indonesia"
	adiva.Age = 25

	fmt.Println(adiva)
	fmt.Println("Nama :", adiva.Name)
	fmt.Println(adiva.Address)
	fmt.Println(adiva.Age)

	// Cara Deklari Struct 2
	udin := Customer{
		Name: "Udin Petot",
		Address: "Jateng",
		Age: 31,
	}
	fmt.Println(udin)

	// Cara Deklari Struct 3
	ucup := Customer{"Ucup Surucup", "Bandung", 32}
	fmt.Println(ucup)
}