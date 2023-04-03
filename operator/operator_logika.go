package main

import "fmt"

func main(){
	username := "petanikode"
	password := "r4h4siaP";

	validUsername := username == "petanikode"
	validPassword := password == "TruePassw012d"

	// operator logika AND
	result := validUsername && validPassword

	fmt.Println(result)
}
