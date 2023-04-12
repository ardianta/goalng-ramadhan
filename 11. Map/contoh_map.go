// membuat map baru
var account map[string]string

// mengisi nilai awal map
account = map[string]string{}

// bisa juga mengisi seperti ini
account = map[string]string{
	"username": "satoshi",
	"pin": "123",
	"saldo":"0"
}

// mengisi data ke map
account["email"] = "dian@petanikode.com"

// mengambil data dari map
fmt.Println(account["email"])

// menggunakan perulangan untuk map
for key, value := range account {
	fmt.Println(key, value)
}

// hapus data di map
delete(account, "email")


// menyimpan maps di dalam slice
var chickens = []map[string]string{
    map[string]string{"name": "chicken blue",   "gender": "male"},
    map[string]string{"name": "chicken red",    "gender": "male"},
    map[string]string{"name": "chicken yellow", "gender": "female"},
}
