package main

import "fmt"

func main() {
	var progress float32 = 4.9333;

	// konversi tipe data ke integer
	progressInteger := uint(progress)

	fmt.Printf("Tipe progress: %T\n", progress)
	fmt.Printf("Nilai progress: %v\n", progress)
	fmt.Printf("Tipe progressInteger: %T\n", progressInteger)
	fmt.Printf("Nilai progressInteger: %v\n", progressInteger)
}
