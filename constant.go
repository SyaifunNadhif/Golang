package main

import "fmt"

func main() {
	// constant --> variabel yang tidak bisa diubah

	const phi float32 = 3.14
	r := 100
	var keliling float32
	keliling = 2 * phi * float32(r)

	fmt.Println("phi :", phi)
	fmt.Println("r   :", r)
	fmt.Println("k   :", keliling)

	const (
		firstName string = "Nadhif"
		lastName         = "Maulana"
		value            = 100
	)

	fmt.Println("\nname     :", firstName, lastName)
	fmt.Println("nilai    :", value)

}
