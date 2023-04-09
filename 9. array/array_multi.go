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
	// +4. [✅] Buat kode buat ganti PIN

func main() {
	var accounts = [3][3]string{
		{"satoshi", "123", "100000"},
		{"dian", "111", "7000000"},
		{"petani", "123", "0"},
	}

	var currentUserIndex int
	
	Auth: {
		var username, pin string
		fmt.Print("Enter Username: ")
		fmt.Scanln(&username)
		fmt.Print("Enter PIN: ")
		fmt.Scanln(&pin)
		fmt.Print("\033[H\033[2J") // clear screen

		for index, account := range(accounts) {
			// cek username and PIN
			if username == account[0] && pin == account[1] {
				currentUserIndex = index
				goto MainMenu
			}
		}
		fmt.Println("⛔ username dan PIN salah")
		goto Auth
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
		fmt.Print("Go to Menu> ")
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
		fmt.Println("Hi", accounts[currentUserIndex][0], "ini sisa saldomu.")
		fmt.Println("💵 Sisa saldo:", accounts[currentUserIndex][2])
		fmt.Scanln()

		goto MainMenu // back to menu
	}

	TopUp: {
		var topUpAmount int
		fmt.Println("💵 Sisa saldo:", accounts[currentUserIndex][2])
		fmt.Print("Top up amount: ")
		fmt.Scanln(&topUpAmount)

		currentBalance, _ := strconv.Atoi(accounts[currentUserIndex][2])
		// top up balance
		currentBalance += topUpAmount

		// update user balance
		accounts[currentUserIndex][2] = strconv.Itoa(currentBalance)

		fmt.Println("✅ Saldo berhasil ditambah")

		goto MainMenu // back to menu
	}

	TarikTunai: {
		var withdrawAmount int
		fmt.Println("💵 Sisa saldo:", accounts[currentUserIndex][2])
		fmt.Print("Jumlah penarikan: ")
		fmt.Scanln(&withdrawAmaount)

		currentBalance, _ := strconv.Atoi(accounts[currentUserIndex][2])
		if currentBalance > withdrawAmount {
			currentBalance -= withdrawAmount
			// update user balance
			accounts[currentUserIndex][2] = strconv.Itoa(currentBalance)
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
		accounts[currentUserIndex][1] = newPIN
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
