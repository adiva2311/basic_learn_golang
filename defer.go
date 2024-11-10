package main

import "fmt"

func logging(){
	fmt.Println("End of Application")
}

func runApplication(){
	defer logging()
	fmt.Println("Running Application")
}



func main(){
	runApplication()
}