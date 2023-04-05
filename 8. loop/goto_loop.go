package main

import "fmt"

func main() {

	count := 0
	goto Ulangi

	Ulangi:
		count++
		fmt.Println("Belajar Golang")
		if count > 10 {
			goto Exit
		} else {
			goto Ulangi
		}


	Exit:
		fmt.Println("Ok done")
}
