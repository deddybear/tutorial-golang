package main

import "fmt"

func main() {
	var c = 0
	const (
		a = 10
		b = 20
	)
	// example sum
	c = a + b
	fmt.Println(c)

	c = b - a
	fmt.Println(c)

	c = b / a
	fmt.Println(c)

	c = b * a
	fmt.Println(c)

	c = b % a
	fmt.Println(c)

	// augmented assignments
	fmt.Println("Augmented Assignments")

	c += a
	fmt.Println(c)
	c -= a
	fmt.Println(c)
	c *= a
	fmt.Println(c)
	c %= a
	fmt.Println(c)

	c++
	fmt.Println(c)
}
