 package main

 import "fmt"

 func main(){

	 samaDengan := 10 == 5
	 tidakSama := 10 != 5
	 lebihBesar := 10 > 5
	 lebihKecil := 10 < 5
	 lebihBesarSama := 10 >= 5
	 lebihKecilSama := 10 <= 5

	 fmt.Println("10 == 5 adalah", samaDengan)
	 fmt.Println("10 != 5 adalah", tidakSama)
	 fmt.Println("10 > 5 adalah", lebihBesar)
	 fmt.Println("10 < 5 adalah", lebihKecil)
	 fmt.Println("10 >= 5 adalah", lebihBesarSama)
	 fmt.Println("10 <= 5 adalah", lebihKecilSama)
}
