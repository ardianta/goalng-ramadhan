package main

import "fmt"

func main(){
	saySomething()
}

func saySomething() {
	fmt.Print("Hello ")
	saySomething()
}
