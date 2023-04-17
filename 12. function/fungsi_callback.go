package main

import "fmt"

func main(){

	userRole := "admin"
	adminFeatureCallback := func(){
		fmt.Println("Fitur Admin")
	}

	isAdmin(userRole, adminFeatureCallback)
}


func isAdmin(role string, next func()){
	if role == "admin" {
		next()
	} else {
		fmt.Println("Akses ditolak, anda bukan admin")
	}

}
