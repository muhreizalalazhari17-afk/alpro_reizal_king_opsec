package main

import "fmt"

func main() {
	var name string

	name = "Yudha Islami Sulistya"
	fmt.Println("Nama : ", name)

	// name = 17
	// fmt.Println(name)

	var lastName = "Sulistya"
	fmt.Println("Nama Belakang : ", lastName)

	middleName := "Islami"
	fmt.Println("Nama Tengah: ", middleName)

	var (
		fullName  = "Yudha Islami Sulistya"
		firstName = "Yudha"
	)

	fmt.Println(fullName)
	fmt.Println(firstName)

}
