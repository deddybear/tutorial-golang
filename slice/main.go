package main

import "fmt"

func main() {

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"November",
		"Oktober",
		"Desember",
	}

	//declarate slice
	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	// var slice2 = months[6:9]

}
