package main

import "fmt"

func main(){
	var a int = 13
	var b int = 5

	operasiTambah:= a + b
	operasiKurang:= a - b
	operasiKali:= a * b
	operasiBagi:= a / b
	operasiModulo:= a % b

	fmt.Println("a =", a, " ", "b =", b)
	fmt.Println("a + b =", operasiTambah)
	fmt.Println("a - b =", operasiKurang)
	fmt.Println("a * b =", operasiKali)
	fmt.Println("a / b =", operasiBagi)
	fmt.Println("a % b =", operasiModulo)
}
