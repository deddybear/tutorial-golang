package main

import "fmt"

func main() {
	type KTP string
	type StatusMarried bool

	var noKTP KTP = "8945796132"
	var status StatusMarried = true

	fmt.Println("No KTP :", noKTP, "Status Menikah :", status)
}
