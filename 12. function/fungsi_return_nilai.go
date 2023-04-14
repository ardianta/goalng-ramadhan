package main

import "fmt"

func main(){
	hasil := jumlahkan(1, 2)
	fmt.Println("Hasil:", hasil)
}

func jumlahkan(nilai1, nilai2 int) (hasil int) {
	hasil = nilai1 + nilai2
	return
}
