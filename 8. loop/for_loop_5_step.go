package main

import "fmt"

func main() {
	for count := 100; count >= 1; count -= 5 {
		fmt.Println("Pengulangan ke", count)
	}
}
