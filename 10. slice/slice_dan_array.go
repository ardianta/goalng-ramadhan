package main

import "fmt"

func main(){
	// membuat array
	names := [3]string{"Linus", "Denis", "Bob"}

	// membuat slice berdasarkan array names
	sliceNames := names[:]

	fmt.Println(names)

	sliceNames[1] = "Dian"

	fmt.Println(names)


	var days = [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// sliceOfDays := days[2:7]

	fmt.Println(days[2:7])
}
