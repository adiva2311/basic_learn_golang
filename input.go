package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&x) // Menggunakan & untuk mengisi nilai x dengan input pengguna
	fmt.Println("Nilai x adalah:", x)
}
