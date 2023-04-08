package main

import (
	"fmt"
)

func main(){
	var ulangi string = "y"
	var count uint = 0

	for ulangi == "y" {
		count++
		fmt.Println("Perulangan ke", count)
		fmt.Println("Mau ulang lagi [y/t]?")
		fmt.Print("Jawab> ")
		fmt.Scanln(&ulangi)
	}

	fmt.Println("Done, sudah diulangi", count, "kali")
}
