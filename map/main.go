package main

import "fmt"

func main() {

	person := map[string]int{
		"age":  1,
		"anak": 3,
	}

	person["tinggi"] = 170

	fmt.Println(person)
	fmt.Println(len(person))
	delete(person, "anak")

	fmt.Println(person)
}
