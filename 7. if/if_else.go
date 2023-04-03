package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	var username, password string

	fmt.Println("** Welcome to Petanikode App **")
	fmt.Print("Username: ")
	if scanner.Scan() {
		username = scanner.Text()
	}

	fmt.Print("Password: ")
	if scanner.Scan() {
		password = scanner.Text()
	}

	// menggunakan if/else
	if username == "petanikode" && password == "petanikode" {
		fmt.Println("MENU:")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. List Data")
		fmt.Println("3. Cari Data")
		fmt.Println("4. Ubah Data")
		fmt.Println("5. Hapus Data")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu> ")
		scanner.Scan()
	} else {
		fmt.Println("⛔ Password dan username salah, coba lagi!")
	}
}
