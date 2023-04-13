package main

import "fmt"

func main(){
	// memanggil fungsi dengan parameter
	ucapkanSalam("Pagi")
	ucapkanSalam("Sore")
	ucapkanSalam("Siang")
	ucapkanSalam("Malam")
	ucapkanSalam("Berpuasa")
}

// membuat fungsi
func ucapkanSalam(konteks string){
	fmt.Println("Selamat", konteks)
}
