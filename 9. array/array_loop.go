var namaArray = [5]int{1,2,3,4,5}

// menggunakan perulangan & array
for i := 0; i < len(namaArray); i++ {
	fmt.Print(namaArray[i])
}
// menggunakan fungsi range
for key, item := range(namaArray) {
	fmt.Println(key, item)
}
