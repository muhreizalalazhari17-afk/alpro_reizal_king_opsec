package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("masukkan suhu")
	fmt.Scanln(&celcius)

	fmt.Println(celcius + 273)
}
