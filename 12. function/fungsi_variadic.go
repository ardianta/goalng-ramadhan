package main

import "fmt"

func main(){
	// memanggil fungsi dengan 5 parameter
	hasil1 := hitungRataRata(1, 2, 1, 3, 4)

	// memanggil fungsi dengan 8 parameter
	hasil2 := hitungRataRata(1, 2, 1, 3, 4, 7, 7, 6)

	fmt.Println("Hasil1 =", hasil1)
	fmt.Println("Hasil2 =", hasil2)
}


func hitungRataRata(numbers ...int) float64 {
    var total int = 0
    for _, number := range numbers {
        total += number
    }

    var avg = float64(total) / float64(len(numbers))
    return avg
}
