package main

import "fmt"

func main(){
	var a int = 10

	fmt.Println("a =",a)
	a += 1
	fmt.Printf("a += 1 = %v\n", a)
	a -= 1
	fmt.Printf("a -= 1 = %v\n", a)
	a *= 5
	fmt.Printf("a *= 5 = %v\n", a)
	a /= 2
	fmt.Printf("a /= 2 = %v\n", a)
	a %= 2
	fmt.Printf("a %%= 2 = %v\n", a)

}
