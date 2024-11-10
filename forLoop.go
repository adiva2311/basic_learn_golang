package main

import "fmt"

func main() {

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	for index, day := range days{
		fmt.Println("Hari ke", index, "-", day)
	}
}