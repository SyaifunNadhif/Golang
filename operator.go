package main

import "fmt"

func main() {
	var a = 5
	var b = 2
	
	if a % 2 == 0{
		fmt.Println(a,"adalah Genap")
	}else{
		fmt.Println(a,"adalah Ganjil")
	}

	fmt.Println("Penjumlahan :", a + b)
	fmt.Println("Pengurangan :", a - b)
	fmt.Println("Perkalian   :", a * b)
	fmt.Println("Pembagian   :", a / b)
	fmt.Println("cek angka   :", a % b)


	a += 10
	fmt.Println("a :", a)
	a--
	fmt.Println("a :", a)



}