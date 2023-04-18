package main

import "fmt"

func main() {
	var pointerScore *int

	if true {
		var score int = 0
		pointerScore = &score;
		*pointerScore = 10
	}

	fmt.Println(score)

}
