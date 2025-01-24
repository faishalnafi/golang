// Program belajar GO-Lang
// Created 13 Maret 2024
package main

import "fmt"

func main() {
	//konversi tipe data1
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//konversi tipe data2
	var nama = "napikk"
	var i byte = nama[0]
	var kata string = string(i)

	fmt.Println(nama)
	fmt.Println(kata)
}
