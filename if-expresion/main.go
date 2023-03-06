package main

import "fmt"

func main() {

	const passed = 80

	const value = 80

	if value > passed {
		fmt.Println("Win")
	} else if value == passed {
		fmt.Println("Draw")
	} else {
		fmt.Println("Lose")
	}

	// short statement
	name := "teeeeest"

	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("ok good")
	}
}
