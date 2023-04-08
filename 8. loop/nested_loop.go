package main

import (
	"fmt"
)

func main(){
	var selectedNumber int
	var gameCount = 0
	var tryAgain string = "y"

	// main loop
	for {
		gameCount++
		fmt.Printf("\n--- Game ke-%v ---\n", gameCount)
		fmt.Print("🎁: ")
		for misteryNumber := 1; misteryNumber <= 5; misteryNumber++ {
			fmt.Print(misteryNumber," ")
		}

		fmt.Println()
		fmt.Print("Pilih angka> ")
		fmt.Scanln(&selectedNumber)

		switch(selectedNumber){
			case 1:
				fmt.Println("Selamat kamu dapat THR: Rp 500.000 🤑")
			case 2:
				fmt.Println("Yay, 🎉 kamu menang tiket liburan ke Rusia")
			case 3:
				fmt.Println("Zonk! Coba lagi 😹")
			case 4:
				fmt.Println("Kamu menang iPhone 14 🤩")
			case 5:
				fmt.Println("Kamu kurang hoki 😜")
			default:
				fmt.Println("Pilih angka yang bener donk!")
		}

		// tanya user, mau main lagi gak?
		fmt.Println("---")
		fmt.Println("Mau main lagi?")
		fmt.Print("Jawab [y/t]> ")
		fmt.Scanln(&tryAgain)

		if(tryAgain != "y"){
			break
		}
	}

	fmt.Println("Done! GAME OVER")
}
