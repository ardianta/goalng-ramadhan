package main

import "fmt"

func main(){

	var a int = 6
	var b int = 3

	operasiBitAnd := a & b
	operasiBitOr := a | b
	operasiBitXor := a ^ b
	operasiBitLeftShift := a << 1
	operasiBitRightShift := a >> 1

	fmt.Println("--- BITWISE AND ---")
	fmt.Printf("%v & %v = %v\n", a, b, operasiBitAnd)
	fmt.Printf("%b & %b = %b\n", a, b, operasiBitAnd)
	fmt.Println("")

	fmt.Println("--- BITWISE OR ---")
	fmt.Printf("%v | %v = %v\n", a, b, operasiBitOr)
	fmt.Printf("%b | %b = %b\n", a, b, operasiBitOr)
	fmt.Println("")

	fmt.Println("--- BITWISE XOR ---")
	fmt.Printf("%v ^ %v = %v\n", a, b, operasiBitXor)
	fmt.Printf("%b ^ %b = %b\n", a, b, operasiBitXor)
	fmt.Println("")

	fmt.Println("--- BITWISE LEFT SHIFT ---")
	fmt.Printf("%v << 1 = %v\n", a, operasiBitLeftShift)
	fmt.Printf("%b << 1 = %b\n", a, operasiBitLeftShift)
	fmt.Println("")

	fmt.Println("--- BITWISE RIGHT SHIFT ---")
	fmt.Printf("%v >> 1 = %v\n", a, operasiBitRightShift)
	fmt.Printf("%b >> 1 = %b\n", a, operasiBitRightShift)
	fmt.Println("")
}
