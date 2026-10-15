package main

import "fmt"

func main() {
	var umur int8
	var suhu float32

	// suhu = 36.3
	// umur = 10
	fmt.Scan(&suhu, &umur)

	fmt.Println("Umur: ", umur)
	fmt.Println("Suhu: ", suhu)
	fmt.Println("Alamat memori dari Var suhu ", &suhu)
	fmt.Println("Alamat memori dari Var umur ", &umur)
}

// 0 0 0 0 0 0 0 0

// 0 = 1 -> 2ˆ0
// 0 = 2 -> 2ˆ1
// 0 = 4 -> 2ˆ2
// 0 = 8 -> 2ˆ3
// 0 = 16 -> 2ˆ4
// 0 = 32 -> 2ˆ5
// 0 = 64 -> 2ˆ6
// 0 = 128 -> 2ˆ7

// Desimal 17 -> 0001 0001
// Desimal 129 -> 1000 0001

// Digital: 1100 1100 -> 204 = 128 + 64 + 8 + 4

// Digital: 0100 0111 -> 71 << 2
// Digital: 0001 1100 -> 28

// Digital: 1110 0101 << 3 -> 229 << 3
// Digital: 0001 0100 -> 20
