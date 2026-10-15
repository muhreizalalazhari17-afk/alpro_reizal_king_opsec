package main

import "fmt"

func main() {
	var y, x int
	fmt.Scan(&y, &x)
	sisa := y % x
	fmt.Println(sisa)
}
