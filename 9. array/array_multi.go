package main

import "fmt"

	// CHALLENGE!
	// 1. [ ] Buat kode untuk menampilkan saldo user
	//    - [ ] input username dan pin, setelah itu tampilkan sisa saldo
	//          dari user yang sedang login (hint: pakai algo searcing)
	// 2. [ ] Buat kode untuk top up saldo
	// 3. [ ] Buat kode untuk tarik saldo
	// +4. [ ] Buat kode buat ganti PIN

func main() {
	var accounts = [3][3]string{
		{"satoshi", "123", "100000"},
		{"dian", "111", "7000000"},
		{"petani", "123", "0"},
	}

	var username, pin string
	var selectedMenu, currentUserIndex int;

	Auth: {
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
		goto Auth
	}

	MainMenu: {
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
				fmt.Println("Menu yang anda pilih salah!")
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
		fmt.Println("TOP UP SALDO")
		goto MainMenu // back to menu
	}

	TarikTunai: {
		fmt.Println("Menu Tarik Tunai")
		goto MainMenu // back to menu
	}

	ChangePIN: {
		fmt.Println("Menu Ubah PIN")
		goto MainMenu // back to menu
	}

	Exit: {
		fmt.Println("Done, Terimakasih!")
	}


}
