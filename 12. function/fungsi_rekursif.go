package main

import "fmt"

func main(){
	hasil := faktorial(9)
	fmt.Println(hasil)
}

func faktorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * faktorial(n-1)
}
