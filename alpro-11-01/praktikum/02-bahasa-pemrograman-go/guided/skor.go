package main

import "fmt"

func main() {
	var nama string
	var matematika, bahasaInggris int
	
	//membaca input dari user
	fmt.Scan(&nama, &matematika, &bahasaInggris)

	//menghitung total dan rata-rata
	total := matematika + bahasaInggris
	rataRata := total / 2

	//menampilkan output dari user
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)
}
