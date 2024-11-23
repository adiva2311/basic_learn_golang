package main

import "fmt"

/*
type NAMA_STRUCT struct{
	Nama_field Type_data
}
*/
type Customer struct{
	Name string
	Address string
	Age int
}

func main(){
	// Cara Membuat Data/Object dari Struct 1
	// var NAMA_VARIABLE NAMA_STRUCT
	var adiva Customer
	adiva.Name = "Adiva Nursuandy"
	adiva.Address = "Indonesia"
	adiva.Age = 25

	fmt.Println(adiva)
	fmt.Println("Nama :", adiva.Name)
	fmt.Println("Alamat :",adiva.Address)
	fmt.Println("Usia :",adiva.Age)

	// Cara Membuat Data/Object dari Struct 2
	udin := Customer{
		Name: "Udin Petot",
		Address: "Jateng",
		Age: 31,
	}
	fmt.Println(udin)

	// Cara Membuat Data/Object dari Struct 3
	ucup := Customer{"Ucup Surucup", "Bandung", 32}
	fmt.Println(ucup)

}