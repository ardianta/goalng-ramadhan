package main

import "fmt"

func main(){
	var score int
	var grade string

	fmt.Println("** Program Grade **")
	fmt.Print("Input score: ")
	fmt.Scanln(&score)

	// menggunakan if/else if/else
	if score >= 80 {
		grade = "A"
	} else if score >= 60 {
		grade = "B"
	} else if score >= 50 {
		grade = "C"
	} else {
		grade = "D"
	}

	fmt.Println("Your Grade: ", grade)
}
