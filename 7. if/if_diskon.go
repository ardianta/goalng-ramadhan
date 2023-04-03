package main

import "fmt"

func main(){
	var totalBelanja int
	var totalToPay int

	var discountAmount float32

	fmt.Println("== Checkout ==")
	fmt.Print("Total Belanja: ")
	fmt.Scan(&totalBelanja)

	if totalBelanja > 100000 {
		fmt.Println("Selamat kamu dapat diskon 30%")
		discountAmount = float32(totalBelanja) * (30.0/100.0)
		totalToPay = totalBelanja - int(discountAmount)
	}

	fmt.Println("Diskon:", discountAmount)

	fmt.Println("Total pembayaran:", totalToPay)
	fmt.Println("---")
	fmt.Println("Terimakasih sudah berbelanja!")
}
