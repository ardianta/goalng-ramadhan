package main

import "fmt"

func main(){
	member := false
	//discount := 5000;
	totalBelanja := 99000

	belanjaDiAtas100k := totalBelanja > 100000

	// Syarat: dia member atau belanjaDiAtas100k
	// operator logika apa yang harus dipakai?
	isDapatDiskon := member || belanjaDiAtas100k

	fmt.Println(isDapatDiskon)
}
