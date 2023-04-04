package main

import "fmt"

func main() {
	var score int
	var grade string

	fmt.Println("## Program Grade ##")
	fmt.Print("Input: ")
	fmt.Scanln(&score)

	switch {
		case score >= 80: {
			grade = "A"
			fmt.Println("Bagus pertahankan!")
			goto HiddenGem
		}
		case score >= 60: {
			grade = "B"
			fmt.Println("Mantap lah, lumayan")
			goto Ending
		}
		case score >= 50: {
			grade = "C"
			fmt.Println("Coba lebih giat lagi nak!")
			goto Ending
		}
		default: {
			grade = "D"
			fmt.Println("Kamu harus remidi")
			fmt.Println("Belajar yang rajin coy!")
			fmt.Println("Jangan menyarah...")
			goto Ending
		}
	}

	HiddenGem: {
		fmt.Println("Hore! Kamu dapat 99 Gem.")
	}

	Ending:
		fmt.Println("Grade anda adalah: ", grade)
}
