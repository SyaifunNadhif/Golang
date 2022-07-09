package main

import "fmt"

func main() {
	var name string
	name = "Maulana"
	// fmt.Println(name)

	name = "Nadhif"
	fmt.Println("nickname :", name)

	var age int8
	age = 20
	fmt.Println("umur     :", age)

	var address = "Demak"
	fmt.Println("alamat   :", address)

	// deklarasi awal variabel
	country := "Indonesia"
	fmt.Println("negara   :", country)

	var (
		firstName = "Nadhif"
		lastName  = "Maulana"
	)

	fmt.Println("name     :", firstName, lastName)
}
