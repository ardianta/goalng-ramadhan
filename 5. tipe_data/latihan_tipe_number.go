package main

import "fmt"

func main(){
	var healthPoint uint8 = 100
	var live uint8 = 3
	var progress float32 = 0.5333333

	progressInPercent := progress * 100

	fmt.Println("== STATUS ==")
	fmt.Println("HP:", healthPoint)
	fmt.Println("Live:", live)
	fmt.Printf("progress: %.2f\n", progress)
	fmt.Printf("Persentase progress: %.2f%%\n", progressInPercent)
}
