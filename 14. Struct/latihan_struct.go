package main

import (
	"fmt"
	"os"
	"bufio"
)

// mendefinisikan struct product
type product struct {
	name string
	price float32
	stock int
}

// slice yang berisi semua object product
var products = []product{}

func main(){
	// main Loop
	for {
		showMainMenu()
	}
}

func showMainMenu(){
	fmt.Println("📦📦📦 Product Management App v.0.0.1 📦📦📦")
	fmt.Println("------------------------------------------")
	fmt.Println("1. Show all products")
	fmt.Println("2. Add new Product")
	fmt.Println("3. Edit Product")
	fmt.Println("4. Delete Product")
	fmt.Println("0. Exit")
	fmt.Print("Choose menu> ")

	var selectedMenu int
	fmt.Scanln(&selectedMenu)
	fmt.Print("\033[H\033[2J") // clear screen
	switch(selectedMenu) {
		case 1:
			showProducts()
		case 2:
			addProduct()
		case 3:
			editProduct()
		case 4:
			deleteProduct()
		case 0:
			exit()
		default:
			fmt.Println("⛔ Invalid selected menu!")
			showMainMenu()
	}
}

func showProducts(){
	if(len(products) <= 0){
		fmt.Println("Product Empty, Please add a new one")
		return
	}

	fmt.Printf("No\tName\tStock\tPrice\n")
	for index, product := range products {
		fmt.Printf("%v\t%v\t%v\t%v\n", index, product.name, product.stock, product.price)
	}
}

func addProduct(){
	scanner := bufio.NewScanner(os.Stdin)
	var name string
	var price float32
	var stock int

	fmt.Print("Product name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Print("Harga: ")
	fmt.Scanln(&price)

	fmt.Print("stok: ")
	fmt.Scanln(&stock)

	var newProduct = product{
		name: name,
		price: float32(price),
		stock: stock,
	}

	products = append(products, newProduct)
	showMainMenu();
}

func editProduct(){
	showProducts()

	fmt.Print("Pilih Produk untuk diedit> ")
	var selectedIndex int
	fmt.Scanln(&selectedIndex)

	if selectedIndex > len(products) || selectedIndex < 0 {
		fmt.Println("⛔ Invalid product index")
		if(len(products) <= 0){
			showMainMenu()
		}
		editProduct()
	}

	scanner := bufio.NewScanner(os.Stdin)
	var name string
	var price float32
	var stock int

	fmt.Print("Product name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Print("Harga: ")
	fmt.Scanln(&price)

	fmt.Print("stok: ")
	fmt.Scanln(&stock)

	var newProduct = product{
		name: name,
		price: float32(price),
		stock: stock,
	}

	// update product data
	products[selectedIndex] = newProduct
}

func deleteProduct(){
	showProducts()

	fmt.Print("Pilih Produk untuk dihapus> ")
	var selectedIndex int
	fmt.Scanln(&selectedIndex)

	if selectedIndex > len(products) || selectedIndex < 0 {
		fmt.Println("⛔ Invalid product index")
		if(len(products) <= 0){
			showMainMenu()
		}
		deleteProduct()
	}

	// delete product
	products[selectedIndex] = products[len(products)-1]
	products = products[:len(products)-1]
	fmt.Println("Done, product deleted!")
	showMainMenu()
}

func exit(){
	os.Exit(0)
}
