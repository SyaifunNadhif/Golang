package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 128
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai32)
	)

	fmt.Println("nilai32 :", nilai32)
	fmt.Println("nilai64 :", nilai64)
	fmt.Println("nilai8  :", nilai8)

	var name = "Nadhif"
	var e byte = name[0]
	var eString = string(e)

	fmt.Println("e       :", e)
	fmt.Println("eString :", eString)
}