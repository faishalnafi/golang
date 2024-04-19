package main

import "fmt"

func main() {
	var tinggi int = 90;
	fmt.Println("value tinggi= ",tinggi)
	fmt.Println("value tinggi= ",&tinggi)
	fmt.Println("value tinggi= ",*&tinggi)

	
}