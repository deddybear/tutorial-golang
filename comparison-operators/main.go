package main

import "fmt"

func main() {

	const (
		number1 = 100
		number2 = 200
	)

	var result = number1 == number2

	fmt.Println(result)

	fmt.Println(number1 == number2)
	fmt.Println(number1 < number2)
	fmt.Println(number1 > number2)
	fmt.Println(number1 != number2)

}
