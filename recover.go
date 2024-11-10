package main

import "fmt"


func endApp(){
	fmt.Println("END APP")
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("Ada Yang Salah")
	}
}

func main(){
	runApp(true)
	fmt.Println("Halo Program Berjalan")
}