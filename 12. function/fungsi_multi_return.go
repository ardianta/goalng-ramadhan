package main

import "fmt"

func main() {
	totalBelanja, persentaseDiskon := 100000, 15
	diskon, totalBayar := calculateDiscount(totalBelanja, persentaseDiskon)

	fmt.Println("belanja:", totalBelanja)
	fmt.Printf("diskon: %v (%v%%)\n", diskon, persentaseDiskon)
	fmt.Println("bayar:", totalBayar)
}

func calculateDiscount(amount, percentage int) (int, int) {
	discount := int(float32(amount) * (float32(percentage)/100.0))
	totalToPay := amount - discount
	return discount, totalToPay
}
