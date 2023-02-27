package main

import "fmt"

func main() {

	// single declarate constant
	const typed string = "value"
	// declarate constant without data type
	const name = "number"
	const value = 80

	fmt.Println(typed, name, value)

	// multiplete declarate constant wiht data type and without data type
	const (
		jenis string = "Penilaian"
		nama         = "Angka"
		nilai        = 80
	)

	fmt.Println(jenis, nama, nilai)
}
