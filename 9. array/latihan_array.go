package main

import "fmt"

func main(){
	// membuat array dengan panjang 5
	var names [5]string

	// mengisi nilai ke array
	names[0] = "Petani"
	names[1] = "Kode"
	names[2] = "Satoshi"
	names[3] = "Sasakmoto"
	names[4] = "Dennis Ritchie"

	// mengecek panjang array
	fmt.Println("Panjang Array:", len(names))

	// mengambil data dari array
	fmt.Println("0.", names[0])
	fmt.Println("1.", names[1])
	fmt.Println("2.", names[2])
	fmt.Println("3.", names[3])
	fmt.Println("4.", names[4])
}
