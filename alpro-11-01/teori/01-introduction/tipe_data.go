package main

import "fmt"

func main() {
	var umur int8
	var anakKe int8
	var saldo int64

	umur = 26
	anakKe = 1
	saldo = 10000000000

	fmt.Println("Umur saya adalah", &umur)
	fmt.Println("Saya adalah anak ke", anakKe)
	fmt.Println("Saldo saya adalah", saldo)

	// Deklarasi
	var namaLengkap string // Pasti sudah ada alamat memorinya
	// Inisialiasi
	var namaDepan string = "Yudha" // Pasti sudah ada alamat memorinya dan sudah ada valuenya
}
