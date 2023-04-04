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

		var selectedMenu string
		// ini if di dalam if
		if scanner.Scan() {
			selectedMenu = scanner.Text()
		}

		// ini switch di dalam if
		switch selectedMenu {
			case "1":
				fmt.Println("Menu 1. Tambah data dipilih")
			case "2":
				fmt.Println("Menu 2. List data dipilih")
			case "3":
				fmt.Println("Menu 3. Cari data dipilih")
			case "4":
				fmt.Println("Menu 4. Ubah data dipilih")
			case "5":
				fmt.Println("Menu 5. Hapus data dipilih")
			case "0":
				fmt.Println("Program Keluar!")
			default:
				fmt.Println("Salah pilih menu")
		}

	} else {
		fmt.Println("⛔ Password dan username salah, coba lagi!")
	}
}

