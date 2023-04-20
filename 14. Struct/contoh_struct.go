
// mendefinisikan struct baru
type account struct {
	username string
	saldo int
	pin int
}

// membuat variabel dengan tipe data dari struct account
var accountBisnis account

// mengakses member struct
accountBisnis.username = "petanikode"
accountBisnis.saldo = 0
accountBisnis.pin = 123
