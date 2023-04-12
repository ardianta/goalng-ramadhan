package main

import (
	"fmt"
	"strconv"
	"os"
)

	// CHALLENGE!
	// 1. [✅] Buat kode untuk menampilkan saldo user
	//    - [✅] input username dan pin, setelah itu tampilkan sisa saldo
	//          dari user yang sedang login (hint: pakai algo searcing)
	// 2. [✅] Buat kode untuk top up saldo
	// 3. [✅] Buat kode untuk tarik saldo
	// 4. [✅] Buat kode buat ganti PIN
	// 4. [✅] Buat fitur registrasi

func main() {
	var accounts = []map[string]string{
		{"username": "satoshi", "pin": "123", "saldo": "100000"},
		{"username": "dian", "pin": "111", "saldo": "7000000"},
		{"username": "petani", "pin": "123", "saldo": "0"},
	}

	var currentUserIndex int

	WelcomeMenu: {
		fmt.Println("----------------------------")
		fmt.Println("🏧 Welcom to GO Virtual ATM ")
		fmt.Println("----------------------------")
		fmt.Println("1. Masuk ke akun")
		fmt.Println("2. Buat akun baru")
		fmt.Println("3. Keluar")
		fmt.Print("Pilh menu> ")
		var selectedMenu int
		fmt.Scanln(&selectedMenu)
		fmt.Print("\033[H\033[2J") // clear screen
		switch(selectedMenu) {
			case 1:
				goto Auth
			case 2:
				goto CreateAccount
			case 0:
				goto Exit
			default:
				fmt.Println("⛔ Menu yang anda pilih salah!")
				goto WelcomeMenu
		}

	}
	
	Auth: {
		var username, pin string
		fmt.Print("Enter Username: ")
		fmt.Scanln(&username)
		fmt.Print("Enter PIN: ")
		fmt.Scanln(&pin)
		fmt.Print("\033[H\033[2J") // clear screen

		for index, account := range(accounts) {
			// cek username and PIN
			if username == account["username"] && pin == account["pin"] {
				currentUserIndex = index
				goto MainMenu
			}
		}
		fmt.Println("⛔ username dan PIN salah")
		goto Auth
	}

	CreateAccount: {
		var username, pin string
		fmt.Print("Enter Username: ")
		fmt.Scanln(&username)
		fmt.Print("Enter PIN: ")
		fmt.Scanln(&pin)
		fmt.Print("\033[H\033[2J") // clear screen

		var newAccount = map[string]string{
			"username": username,
			"pin": pin,
			"saldo": "0",
		}

		accounts = append(accounts, newAccount)

		fmt.Println("✅ Akun berhasil dibuat")
		fmt.Scanln()
		goto WelcomeMenu
	}

	MainMenu: {
		var selectedMenu int
		fmt.Println("------------------")
		fmt.Println("🏧 GO Virtual ATM ")
		fmt.Println("------------------")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Top up")
		fmt.Println("3. Tarik tunai")
		fmt.Println("4. Ubah PIN")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih Menu> ")
		fmt.Scanln(&selectedMenu)
		fmt.Print("\033[H\033[2J") // clear screen
		switch(selectedMenu) {
			case 1:
				goto CekSaldo
			case 2:
				goto TopUp
			case 3:
				goto TarikTunai
			case 4:
				goto ChangePIN
			case 0:
				goto Exit
			default:
				fmt.Println("⛔ Menu yang anda pilih salah!")
				goto MainMenu
		}
	}

	CekSaldo: {
		// Tampilkan sisa saldo
		fmt.Println("Hi", accounts[currentUserIndex]["username"], "ini sisa saldomu.")
		fmt.Println("💵 Sisa saldo:", accounts[currentUserIndex]["saldo"])
		fmt.Scanln()

		goto MainMenu // back to menu
	}

	TopUp: {
		var topUpAmount int
		fmt.Println("💵 Sisa saldo:", accounts[currentUserIndex]["saldo"])
		fmt.Print("Top up amount: ")
		fmt.Scanln(&topUpAmount)

		currentBalance, _ := strconv.Atoi(accounts[currentUserIndex]["saldo"])
		// top up balance
		currentBalance += topUpAmount

		// update user balance
		accounts[currentUserIndex]["saldo"] = strconv.Itoa(currentBalance)

		fmt.Println("✅ Saldo berhasil ditambah")

		goto MainMenu // back to menu
	}

	TarikTunai: {
		var withdrawAmount int
		fmt.Println("💵 Sisa saldo:", accounts[currentUserIndex]["saldo"])
		fmt.Print("Jumlah penarikan: ")
		fmt.Scanln(&withdrawAmount)

		currentBalance, _ := strconv.Atoi(accounts[currentUserIndex]["saldo"])
		if currentBalance > withdrawAmount {
			currentBalance -= withdrawAmount
			// update user balance
			accounts[currentUserIndex]["saldo"] = strconv.Itoa(currentBalance)
			fmt.Println("✅ Penarikan berhasil!")
		} else {
			fmt.Println("⛔ Saldo tidak cukup!")
		}

		goto MainMenu // back to menu
	}

	ChangePIN: {
		var newPIN string
		fmt.Print("PIN baru: ")
		fmt.Scanln(&newPIN)
		accounts[currentUserIndex]["pin"] = newPIN
		fmt.Println("✅ PIN sudah diubah")
		// Note: pin akan di reset ulang setelah program ditutup
		// untuk improvement, simpan data ke dalam file atau database

		goto MainMenu // back to menu
	}

	Exit: {
		fmt.Println("Done, Terimakasih! 👋 bye bye.")
		os.Exit(0)
	}
}
