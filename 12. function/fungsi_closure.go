package main

import "fmt"

func main(){
	// contoh fungsi closure
	var ucapSalam = func(konteks string){
		fmt.Println("Selamat", konteks)
	}

	ucapSalam("Pagi")
	fmt.Printf("%T\n", ucapSalam)


	// Immediately-Invoked Function Expression (IIFE)
	func(ctx string){
		fmt.Println("Selamat", ctx)
	}("Berpuasa")


	func(){
		fmt.Println("Yoo! saya fungsi closure")
		fmt.Println("yang dipanggil dengan IIFE")
	}()
}
