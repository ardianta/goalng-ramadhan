package main

import "fmt"

func main() {
	var score int = 0;
	addScore(&score, 10)
	fmt.Println("Isi var score di main:", score)
}

func addScore(score *int, newScore int) {
	*score += newScore
	fmt.Println("Isi var score setelah ditambahkan:", *score)
}
