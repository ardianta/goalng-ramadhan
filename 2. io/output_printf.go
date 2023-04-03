package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "=== PROGRAM GO ===")
	fmt.Printf("v%d.%d.%d by %s\n\n", 0,0,1, "Petani Kode")

	fmt.Printf("Menu Sahur\tHarga\n")
	fmt.Println("---------------------")
	fmt.Printf("%v\t%d\n", "1. Nasi Balap", 15000)
	fmt.Printf("%v\t%d\n", "2. Es Teh", 5000)
	fmt.Printf("%v\t%d\n", "3. Sate Golang", 18000)
	fmt.Println("---------------------")
	fmt.Printf("%c Selamat Menikmati %c\n", '🙏','🙏')
}
