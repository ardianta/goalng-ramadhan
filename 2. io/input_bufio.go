package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var name string
	var age int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("### Welcome to Underground Club ###")
	fmt.Print("Enter your name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}


	fmt.Print("Enter your age: ")
	if scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			age = i
		}
	}

	fmt.Println("---------")
	fmt.Printf("Hi %v Welcome to the club!\n", name)
	fmt.Println("You are", age, "years old")
}
