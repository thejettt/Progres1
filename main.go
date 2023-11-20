package main

import (
	"fmt"
	"time"
)

type Pengeluaran struct {
	Name   string
	Date   time.Time
	Amount float64
}

var (
	pengeluaran []Pengeluaran
	loggedIn    bool
)

func main() {
	var choice int

	for {
		fmt.Println("\nMenu :")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Logout")
		fmt.Println("4. Tambah Pengeluaran")
		fmt.Println("5. Tampilkan Seluruh Data")
		fmt.Println("6. Exit")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			login()
		case 2:
			register()
		case 3:
			logout()
		case 4:
			if !loggedIn {
				fmt.Println("Anda perlu login untuk menambahkan pengeluaran.")
				break
			}
			addPengeluaran()
		case 5:
			if !loggedIn {
				fmt.Println("Anda perlu login untuk menampilkan data.")
				break
			}
			menampilkanData()
		case 6:
			fmt.Println("Terima kasih! Sampai jumpa lagi.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
func login() {
	fmt.Println("Login Berhasil!")
	loggedIn = true
}

func register() {
	// Implementasi register
	fmt.Println("Akun berhasil didaftarkan!")
}

func logout() {
	// Implementasi logout
	fmt.Println("Logout Berhasil.")
	loggedIn = false
}
func addPengeluaran() {
	var peng Pengeluaran

	fmt.Println("Masukkan rincian pengeluaran")
	fmt.Print("Nama Pengeluaran :")
	fmt.Scanln(&peng.Name)
	fmt.Print("Tanggal Pengeluaran (dd-mm-yyyy) :")
	dateInput := ""
	fmt.Scanln(&dateInput)
	parsedDate, err := time.Parse("02-01-2006", dateInput)
	if err != nil {
		fmt.Println("format tanggal salah, Gunakan format dd-mm-yyyy")
		return
	}
	peng.Date = parsedDate

	fmt.Print("Jumlah Pengeluaran :")
	fmt.Scanln(&peng.Amount)
	pengeluaran = append(pengeluaran, peng)
	fmt.Println("Pengeluaran berhasil di tambahkan")
}
func menampilkanData() {
	fmt.Println("\nData keseluruhan pengeluaran:")
	for _, p := range pengeluaran {
		fmt.Printf("Nama: %s | Tanggal: %s | Jumlah: %.2f\n", p.Name, p.Date.Format("02-01-2006"), p.Amount)
	}
}
