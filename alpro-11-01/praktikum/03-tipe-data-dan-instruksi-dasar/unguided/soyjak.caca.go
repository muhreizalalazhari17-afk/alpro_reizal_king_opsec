package main

import "fmt"

func main() {
	var totalHari int
	fmt.Scan(&totalHari)

	var tahun int = totalHari / 360
	var sisa int = totalHari % 360

	bulan := sisa / 30
	sisa = sisa % 30

	minggu := sisa / 7
	hari := sisa % 7

	fmt.Println(tahun)
	fmt.Println(bulan)
	fmt.Println(minggu)
	fmt.Println(hari)
}
