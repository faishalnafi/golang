package main

import "fmt"

func main() {
	var tinggi int = 90
	fmt.Println("value tinggi : ", tinggi)
	fmt.Println("alamat tinggi : ", &tinggi)
	fmt.Println("dereference tinggi : ", *&tinggi)
	fmt.Println()

	var alamat *int
	alamat = &tinggi
	fmt.Println("isi alamat : ", alamat)
	fmt.Println("value dari var alamat :", *alamat)
	fmt.Println()

	*alamat = 1000
	fmt.Println("value tinggi :", tinggi)
}