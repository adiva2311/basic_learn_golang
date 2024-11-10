package main

import "fmt"

func endApp(){
	fmt.Println("END APP")
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("Ada Yang Salah")
	}
}

func main(){
	runApp(true)
}