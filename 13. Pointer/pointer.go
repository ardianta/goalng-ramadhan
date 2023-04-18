package main

import "fmt"

func main() {
	var name string = "Petani Kode"
	var balance int = 0

	// membaut pointer
	var pName *string = &name;
	var pBalance *int = &balance;

	fmt.Println("Isi name =", name)
	fmt.Println("Isi balance =", balance)
	fmt.Println("Alamat name:", &name)
	fmt.Println("Alamat balance:", &balance)
	fmt.Println("--- Pointer ---")
	fmt.Println("Alamat *pName:", pName)
	fmt.Println("Alamat *pBalance:", pBalance)

	// mengubah isi variabel dari pointer
	*pName = "Dian"
	*pBalance = 99000

	fmt.Println("Isi name =", name)
	fmt.Println("Isi balance =", balance)
}
