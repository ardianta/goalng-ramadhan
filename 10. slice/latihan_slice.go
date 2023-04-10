package main

import "fmt"

func main(){
	// membuat slice kosong
	var names []string

	// mengisi slice
	names = []string{"Satoshi", "Nakamoto", "Linus", "Torvalds"}

	fmt.Println("isi slice:", names)
	fmt.Println("Panajang:", len(names))
	fmt.Println("Kapasitas:", cap(names))

	// menambah data ke Slice dengan fungsi append()
	names = append(names, "Petani Kode")

	// cek panjang kapasitas Slice setelah ditambahkan
	fmt.Println("Panajang setelah ditambah:", len(names))
	fmt.Println("Kapasitas setelah ditambah:", cap(names))

	// tampilkan lagi isi Slice
	fmt.Println("isi slice setelah ditambah:", names)
}
