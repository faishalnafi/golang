package main

import (
	"fmt"
)

// Struktur data
type Kelompok struct {
	NamaDepan    string
	NamaBelakang string
	Email        string
	Username     string
	Umur         int
}

// Fungsi untuk input data
func dataMasuk(data *Kelompok) {
	fmt.Scanf("Nama depan: %s\n", &data.NamaDepan)
	fmt.Scanf("Nama belakang: %s\n", &data.NamaBelakang)
	fmt.Scanf("Surel: %s\n", &data.Email)
	fmt.Scanf("Nama pengguna: %s\n", &data.Username)
	fmt.Scanf("Usia: %d\n", &data.Umur)
}

func main() {
	var data Kelompok

	// Looping input data
	var jumlahData int
	fmt.Scanf("Jumlah data yang ingin diinput: %d\n", &jumlahData)

	for i := 0; i < jumlahData; i++ {
		fmt.Printf("\nData Id ke-%d\n", i+1)
		dataMasuk(&data)
	}

	// Menampilkan data
	fmt.Println("\n====================")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Id %d\n", i+1)
		fmt.Printf("Nama Lengkap: %s %s\n", data.NamaDepan, data.NamaBelakang)
		fmt.Println("Surel  :", data.Email)
		fmt.Println("Username :", data.Username)
		fmt.Println("Usia  :", data.Umur)
		fmt.Println()
	}
}
