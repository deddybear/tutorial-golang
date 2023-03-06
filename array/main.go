package main

import "fmt"

func main() {
	var data [3]string
	//or
	var value = [3]int{
		80,
		75,
		90,
	}
	//or
	var test = [...]int{
		10,
		22,
		33,
		44,
	}

	data[0] = "John"
	data[1] = "Smith"
	data[2] = "Doe"

	if len(data) == len(value) {
		for i := 0; i < len(value); i++ {
			fmt.Println(data[i], value[i])
		}
	}

	fmt.Println(len(test))
}
