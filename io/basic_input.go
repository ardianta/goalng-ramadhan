package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("### Welcome to Underground Club ###")
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Println("---------")
	fmt.Printf("Hi %v Welcome to the club!\n", name)
	fmt.Println("You are", age, "years old")
}
