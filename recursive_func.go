// Faktorial Recursive adalah dimana Function memanggil Function itu sendiri

package main

import "fmt"

func deretAritmatika (value int) int{
	if value == 1{
		return 1
	} else {
		return value + deretAritmatika(value-1)
	}
}

func factorial (value2 int) int{
	if value2 == 1{
		return 1
	} else {
		return value2 * factorial(value2-1)
	}
}



func main(){
	result := 20+19+18+17+16+15+14+13+12+11+10+9+8+7+6+5+4+3+2+1
	fmt.Println(result)
	fmt.Println(deretAritmatika(20))

	result2 := 5*4*3*2*1
	fmt.Println(result2)
	fmt.Println(factorial(5))
}