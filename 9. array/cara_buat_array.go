// 1. Membuat array kosong dengan panjang 3 dan tipe data string
var nama_array = [3]string{}

// atau bisa gini juga
var nama_array [3]string

// 2. Membuat array dengan panjang 3 dan langsung diisi data
var nama_array = [3]string{
	"Data 1",
	"Data 2",
	"Data 3",
}

// 3. Membuat array beserta isinya tanpa menentukan panjang
var nama_array = [...]string{
	"Data 1",
	"Data 2",
	"Data 3"
	"Data 4"
}

/* Note: Begitu array di buat, panjangnya akan fixed
 * Dia tidak dinamis seperti list di Python
 * Kalau mau buat yang lebih fleksibel, pakai slice aja
 */
