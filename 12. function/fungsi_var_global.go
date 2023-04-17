package main

import "fmt"

appName := "Petani Kode App"
appVersion := "0.0.1"

func main() {
	showAppName()
	fmt.Println("Version: ", appVersion)
}

func showAppName(){
	fmt.Println("===", appName, "===")
}
