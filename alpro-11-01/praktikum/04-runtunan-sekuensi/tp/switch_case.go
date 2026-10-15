package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan nilai (0-100): ")
	fmt.Scan(&nilai)

	switch {
	case nilai >= 80 && nilai <= 100:
		fmt.Println("Indeks: A")
	case nilai >= 70:
		fmt.Println("Indeks: B")
	case nilai >= 60:
		fmt.Println("Indeks: C")
	case nilai >= 0:
		fmt.Println("Indeks: D")
	default:
		fmt.Println("Nilai tidak valid")
	}
}
