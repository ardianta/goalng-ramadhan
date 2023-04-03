package main

import "fmt"

func main() {
	// Membuat variabel baru
	var name = "Satoshi"
	var balance uint = 0
	var topUpAmount uint = 0

	// topup balance
	fmt.Print("Enter topup amount: ")
	fmt.Scanln(&topUpAmount)
	balance += topUpAmount

	// print current balance
	fmt.Println("Hi", name, "👋")
	fmt.Println("Your current balance:", balance)
}
